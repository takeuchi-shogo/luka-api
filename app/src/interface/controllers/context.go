package controllers

type Context interface {
	GetHeader(key string) string
	Header(key, value string)
	JSON(code int, obj interface{})
	Param(key string) string
	PostForm(key string) (value string)
	Query(key string) (value string)
}
