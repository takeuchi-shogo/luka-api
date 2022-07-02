package controllers

type Context interface {
	JSON(code int, obj interface{})
	Param(key string) string
	PostForm(key string) (value string)
}
