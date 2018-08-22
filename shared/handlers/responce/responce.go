package responce

import (
	"net/http"
	"encoding/json"
)

type JsonResponse struct {
	Status string `json:status`
	Code int `json:code`
	Data interface{} `json:data`
}

func (r *JsonResponse) Success(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	r.Status = "success"
	r.Code = http.StatusOK
	r.Data = data

	js, err := json.Marshal(r)
	if err != nil {
		panic(err)
		return
	}
	w.Write(js)
}

func (r *JsonResponse) Error(w http.ResponseWriter, code int, text string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	r.Status = "error"
	r.Code = code
	r.Data = text

	js, err := json.Marshal(r)
	if err != nil {
		panic(err)
		return
	}
	w.Write(js)
}