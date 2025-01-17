package route

var (
	CtxKeyProxyName    = contextKey("proxy name")
	CtxKeyProviderName = contextKey("provider name")
	CtxKeyProxy        = contextKey("proxy")
	CtxKeyProvider     = contextKey("provider")
)

type contextKey string

func (c contextKey) String() string {
	return "mat context key " + string(c)
}
