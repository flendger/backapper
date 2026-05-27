package healthcheckcontroller

import (
	"backapper/app"
	"backapper/app/appholder"
	"backapper/app/appservice"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthCheckController_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := appservice.New(appholder.New(
		&app.App{Name: "ok", HealthCheck: "echo healthy"},
		&app.App{Name: "fail", HealthCheck: "exit 1"},
		&app.App{Name: "no-config"},
	), log.Default())
	controller := New(service)

	tests := []struct {
		name           string
		query          string
		wantStatusCode int
		wantBody       string
	}{
		{
			name:           "missing app param",
			query:          "",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Bad request: no App param\n",
		},
		{
			name:           "unknown app",
			query:          "app=missing",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Couldn't healthcheck app [missing]: app [missing] doesn't exist\n",
		},
		{
			name:           "healthcheck not configured",
			query:          "app=no-config",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Couldn't healthcheck app [no-config]: healthcheck not configured for app [no-config]\n",
		},
		{
			name:           "script exit non-zero",
			query:          "app=fail",
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       "Couldn't healthcheck app [fail]:",
		},
		{
			name:           "success",
			query:          "app=ok",
			wantStatusCode: http.StatusOK,
			wantBody:       "healthy\n\nOK healthcheck: ok\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Request = httptest.NewRequest(http.MethodGet, "/healthcheck?"+tt.query, nil)

			controller.Handle(context)

			if recorder.Code != tt.wantStatusCode {
				t.Fatalf("status = %d, want %d", recorder.Code, tt.wantStatusCode)
			}
			if !strings.Contains(recorder.Body.String(), tt.wantBody) {
				t.Fatalf("body = %q, want to contain %q", recorder.Body.String(), tt.wantBody)
			}
		})
	}
}
