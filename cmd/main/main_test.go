package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/ruelephant/MailRuTestApp/shared/handlers"
	"github.com/ruelephant/MailRuTestApp/shared/handlers/responce"
	"github.com/ruelephant/MailRuTestApp/shared/dao_mock"
	"net/url"
	"github.com/ruelephant/MailRuTestApp/shared/handlers/limiter"
	"github.com/ulule/limiter/drivers/middleware/stdlib"
	"encoding/json"
)

func TestEventAddValidationHandler(t *testing.T) {
	dao := dao_mock.NewEventDaoMock()
	rr, err := NewAddEventRequest(dao,nil,"Z")
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestEventHandler(t *testing.T) {
	dao := dao_mock.NewEventDaoMock()

	rr, err := NewAddEventRequest(dao,nil,"A")
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	rr, err = NewStatRequest(dao,nil)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responcejson *ResponceStat
	json.Unmarshal(rr.Body.Bytes(), &responcejson)

	if responcejson.Data.A != 1 {
		t.Error("A != 1",)
	}
}

func TestLimitHandler(t *testing.T) {
	dao := dao_mock.NewEventDaoMock()

	limiter, err := limiter.GetLimitMiddleware()
	if err != nil {
		t.Fatal(err)
	}

	// Correct request
	rr, err := NewAddEventRequest(dao, limiter,"A")
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Overload
	for i:=0;i<10;i++ {
		rr, err := NewAddEventRequest(dao, limiter,"A")
		if err != nil {
			t.Fatal(err)
		}

		// Check the status code is what we expect.
		if status := rr.Code; status != 509 {
			t.Fatalf("handler returned wrong status code: got %v want %v",
				status, 509)
		}
	}
}


// ------------------ Fake HTTP requests wrappers -------------------
func NewAddEventRequest(dao handlers.EventDao, limiter *stdlib.Middleware, eventType string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("POST", "/events", nil)
	req.RemoteAddr = "127.0.0.1"
	req.PostForm = url.Values{"type": []string{eventType}}
	if err != nil {
		return nil, err
	}

	response := &responce.JsonResponse{}
	rr := httptest.NewRecorder()

	eventHandler := handlers.NewEventHandler(dao, response)
	if limiter != nil {
		handler := limiter.Handler(eventHandler)
		handler.ServeHTTP(rr, req)
	} else {
		handler := eventHandler
		handler.ServeHTTP(rr, req)
	}

	return rr, nil
}

func NewStatRequest(dao handlers.EventDao, limiter *stdlib.Middleware) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("GET", "/stat", nil)
	req.RemoteAddr = "127.0.0.1"
	if err != nil {
		return nil, err
	}

	response := &responce.JsonResponse{}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	statHandler := handlers.NewStatHandler(dao, response)
	if limiter != nil {
		handler := limiter.Handler(statHandler)
		handler.ServeHTTP(rr, req)
	} else {
		handler := statHandler
		handler.ServeHTTP(rr, req)
	}

	return rr, nil
}

// ------------------ END HTTP requests wrappers -------------------

type ResponceStat struct {
	Code int `json:"Code"`
	Data struct {
		A int `json:"A"`
		B int `json:"B"`
		C int `json:"C"`
	} `json:"Data"`
	Status string `json:"Status"`
}