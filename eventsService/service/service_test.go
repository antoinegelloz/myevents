package service

import (
	"bytes"
	"github.com/agelloz/reach/eventsService/configuration"
	"github.com/agelloz/reach/eventsService/persistence"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimple_API_Usage(t *testing.T) {
	dbh, _ := persistence.NewPersistenceLayer(configuration.DBTypeDefault, configuration.DBConnectionDefault)
	var h = &EventsServiceHandler{
		dbHandler:   dbh,
		endpoint:    configuration.EndpointDefault,
		tlsEndpoint: configuration.TLSEndpointDefault,
	}

	t.Run("Get all events", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events", nil)
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.getAllEventsHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Add event by name circle test", func(t *testing.T) {
		var jsonStr = []byte(`{"name":"circle test"}`)
		req, err := http.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(jsonStr))
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.addEventHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Get event by name circle test", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events", nil)
		req = mux.SetURLVars(req, map[string]string{"nameOrID": "name", "nameOrIDValue": "circle test"})
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.getEventHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Delete event by name circle test", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/events", nil)
		req = mux.SetURLVars(req, map[string]string{"nameOrID": "name", "nameOrIDValue": "circle test"})
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.deleteEventHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Get event by name circle test", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events", nil)
		req = mux.SetURLVars(req, map[string]string{"nameOrID": "name", "nameOrIDValue": "circle test"})
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		h.getEventHandler(w, req)
		resp := w.Result()
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
