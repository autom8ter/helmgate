package client

type ctxKey string

const (
	authCtxKey   ctxKey = "x-graphik-auth-ctx"
	methodCtxKey ctxKey = "x-graphik-full-method"
	tokenCtxKey  ctxKey = "x-graphik-token"
)
