package handlers

import (
	"encoding/json"
	"net/http"
)

type eventHttpHandler struct {
	DAO EventDao
}

func (h *eventHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var result map[string]string
	result["status"] = "success"

	js, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func NewEventHandler(dao EventDao) http.Handler {
	return &eventHttpHandler{
		DAO: dao,
	}
}