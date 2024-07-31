package frouter

type Router map[string]*Route

const (
	FieldMetaKey  = "Meta"
	PathMetaKey   = "path"
	MethodMetaKey = "method"
	SignMetaKey   = "sign"
	LoginMetaKey  = "login"
)
