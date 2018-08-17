package handlers

import "net/http"

type statHandler struct {
	DAO EventDao
}

func (h *statHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func NewStatHandler(dao EventDao) *statHandler {
	return &statHandler{
		DAO: dao,
	}
}