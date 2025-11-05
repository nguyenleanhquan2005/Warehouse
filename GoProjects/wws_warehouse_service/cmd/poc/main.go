package main

import (
	"os"
	ocios "app/pkg/oci/os"
	"context"
)

func main() {

	ocios.Setup(false, "/Users/thang.le/.oci/config", "tanthangsport")

	bucketName := "wws-product-images"
	objectName := "sample-image.png"
	filePath := "./assets/screen1.png"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()	
	if err != nil {
		panic(err)
	}

	bucketManager := ocios.GetInstance()
	err = bucketManager.PutObject(context.Background(), bucketName, objectName, fileInfo.Size(), file, nil)
	if err != nil {
		panic(err)
	}
}
