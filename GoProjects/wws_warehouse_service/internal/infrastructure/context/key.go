package context

type key string

func (k key) String() string {
	return "middleware.context.key." + string(k)
}

const (
	keyVendorID key = "vendor_id"
)
