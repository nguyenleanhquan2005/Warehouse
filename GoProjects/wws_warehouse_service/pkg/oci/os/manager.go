package ocios

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
)

type OCIOSManager interface {
}

type osmanager struct {
	Namespace           string
	ObjectStorageClient objectstorage.ObjectStorageClient
}

var (
	instance *osmanager
	once     sync.Once
)

func Setup(useInstancePrinciple bool, configFilePath, profile string) {
	once.Do(func() {
		var provider common.ConfigurationProvider
		var err error
		if useInstancePrinciple {
			provider, err = auth.InstancePrincipalConfigurationProvider()
			if err != nil {
				log.Fatalf("failed to create instance principal provider: %v", err)
			}
		} else {
			provider, err = common.ConfigurationProviderFromFileWithProfile(
				configFilePath, // path to config file
				profile,        // profile name
				"",             // passphrase (if your key is encrypted)
			)
			if err != nil {
				panic(fmt.Errorf("failed to load OCI config: %w", err))
			}
		}

		client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(provider)
		if err != nil {
			panic(fmt.Errorf("failed to create object storage client: %w", err))
		}

		instance = &osmanager{
			ObjectStorageClient: client,
		}

		namespace, err := instance.getNamespace(context.Background(), client)
		if err != nil {
			panic(fmt.Errorf("failed to get namespace: %w", err))
		}

		instance.Namespace = namespace
	})

}

func GetInstance() *osmanager {
	if instance == nil {
		log.Println("Please call Setup function before using GetInstance")
		return nil
	}
	return instance
}

func (m *osmanager) PutObject(ctx context.Context, bucketname, objectname string, contentLen int64, content io.ReadCloser, metadata map[string]string) error {
	request := objectstorage.PutObjectRequest{
		NamespaceName: &m.Namespace,
		BucketName:    &bucketname,
		ObjectName:    &objectname,
		ContentLength: &contentLen,
		PutObjectBody: content,
		OpcMeta:       metadata,
	}
	_, err := m.ObjectStorageClient.PutObject(ctx, request)

	return err
}

func (m *osmanager) getNamespace(ctx context.Context, c objectstorage.ObjectStorageClient) (string, error) {
	request := objectstorage.GetNamespaceRequest{}
	r, err := c.GetNamespace(ctx, request)
	if err != nil {
		return "", err
	}

	return *r.Value, nil
}
