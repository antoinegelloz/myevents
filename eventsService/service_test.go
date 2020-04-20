package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agelloz/reach/eventsService/configuration"
	"github.com/agelloz/reach/eventsService/persistence"
	"github.com/agelloz/reach/eventsService/service"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestSimple_API_Usage(t *testing.T) {
	dbh, _ := persistence.NewPersistenceLayer(configuration.DBTypeDefault, configuration.DBConnectionDefault)
	conn, err := amqp.Dial(configuration.AMPQURLDefault)
	if err != nil {
		panic(err)
	}
	var h = &service.EventsServiceHandler{
		DbHandler:      dbh,
		Endpoint:       configuration.EndpointDefault,
		TLSEndpoint:    configuration.TLSEndpointDefault,
		AMQPConnection: conn,
	}

	t.Run("Get all events", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events", nil)
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.GetAllEventsHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Get event by name circle test", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events", nil)
		req = mux.SetURLVars(req, map[string]string{"nameOrID": "name", "nameOrIDValue": "circle test"})
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.GetEventHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("Add event by name circle test", func(t *testing.T) {
		var jsonStr = []byte(`{"name":"circle test"}`)
		req, err := http.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(jsonStr))
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.AddEventHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Get event by name circle test", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events", nil)
		req = mux.SetURLVars(req, map[string]string{"nameOrID": "name", "nameOrIDValue": "circle test"})
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.GetEventHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Delete event by name circle test", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/events", nil)
		req = mux.SetURLVars(req, map[string]string{"nameOrID": "name", "nameOrIDValue": "circle test"})
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.DeleteEventHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Get event by name circle test", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events", nil)
		req = mux.SetURLVars(req, map[string]string{"nameOrID": "name", "nameOrIDValue": "circle test"})
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.GetEventHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
