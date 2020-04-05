package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8181/events", nil)
	resp := req.Response
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	//fmt.Println(resp.Body)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
