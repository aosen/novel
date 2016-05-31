package controllers

type NotFoundHandler struct {
	BaseHandler
}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}
