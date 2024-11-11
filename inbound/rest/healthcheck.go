package rest

import (
	"net/http"
)

type HealthCheckHandler struct {
	healthChecker HealthChecker
}

func NewHealthCheck(healthChecker HealthChecker) *HealthCheckHandler {
	return &HealthCheckHandler{
		healthChecker: healthChecker,
	}
}

func (h *HealthCheckHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	err := h.healthChecker.Check()
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	_, err = writer.Write([]byte(`{"status": "OK", "errors": []}`))
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}
}
