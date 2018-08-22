package handlers

import "net/http"

type EventDao interface {
	AddEvent(eventType string) error
	CountByType(eventType string) (int, error)
}

type ResponseInterface interface {
	Success(w http.ResponseWriter, data interface{})
	Error(w http.ResponseWriter, code int, text string)
}