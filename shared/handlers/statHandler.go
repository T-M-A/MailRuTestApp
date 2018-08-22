package handlers

import (
	"net/http"
)

type statHandler struct {
	dao EventDao
	response ResponseInterface
}

func (h *statHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]int)

	groups := []string{"A", "B", "C"}

	for _, group := range groups {
		aCount, err := h.dao.CountByType(group)
		if err != nil {
			h.response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		result[group] = aCount
	}

	h.response.Success(w, result)
}

func NewStatHandler(dao EventDao, response ResponseInterface) *statHandler {
	return &statHandler{
		dao: dao,
		response: response,
	}
}