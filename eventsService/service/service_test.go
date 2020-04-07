package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimple_API_Usage(t *testing.T) {
	eh := NewEventsServiceHandler()

	req1, err1 := http.NewRequest(http.MethodGet, "/events", nil)
	if err1 != nil {
		t.Fatal(err1)
	}
	rr1 := httptest.NewRecorder()
	h1 := http.HandlerFunc(eh.getAllEventsHandler)
	h1.ServeHTTP(rr1, req1)
	assert.Equal(t, http.StatusOK, rr1.Code)
	fmt.Print(rr1.Body.String())
/*
	req2, err2 := http.NewRequest("POST", "/events", bytes.NewBuffer([]byte(`{"name":"test"}`)))
	if err2 != nil {
		t.Fatal(err2)
	}
	rr2 := httptest.NewRecorder()
	h2 := http.HandlerFunc(eh.addEventHandler)
	h2.ServeHTTP(rr2, req2)
	assert.Equal(t, http.StatusOK, rr2.Code)
*/
	req3, err3 := http.NewRequest(http.MethodGet, "/events/name/test", nil)
	if err3 != nil {
		t.Fatal(err3)
	}
	rr3 := httptest.NewRecorder()
	h3 := http.HandlerFunc(eh.getEventHandler)
	h3.ServeHTTP(rr3, req3)
	assert.Equal(t, http.StatusOK, rr3.Code)
	fmt.Print(rr3.Body.String())
}
