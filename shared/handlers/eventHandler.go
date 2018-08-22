package handlers

import (
	"net/http"
)

type eventHttpHandler struct {
	dao EventDao
	response ResponseInterface
}

func (h *eventHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	eventType := r.Form.Get("type")
	switch eventType {
		case "A", "B", "С":
			if err != nil {
				h.response.Error(w, http.StatusInternalServerError, err.Error())
				return
			}
			h.dao.AddEvent(eventType)
			h.response.Success(w, nil)
		default:
			h.response.Error(w, http.StatusBadRequest, "Invalid type")
			return
	}
}

func NewEventHandler(dao EventDao, responce ResponseInterface) *eventHttpHandler {
	return &eventHttpHandler{
		dao: dao,
		response: responce,
	}
}