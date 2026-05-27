package restartcontroller

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

func TestRestartController_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := appservice.New(appholder.New(
		&app.App{Name: "ok", Restart: "echo restarted"},
		&app.App{Name: "fail", Restart: "exit 1"},
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
			wantStatusCode: http.StatusOK,
			wantBody:       "Bad request: no App param\n",
		},
		{
			name:           "unknown app",
			query:          "app=missing",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Couldn't restart app [missing]: app [missing] doesn't exist\n",
		},
		{
			name:           "script exit non-zero",
			query:          "app=fail",
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       "Couldn't restart app [fail]:",
		},
		{
			name:           "success",
			query:          "app=ok",
			wantStatusCode: http.StatusOK,
			wantBody:       "restarted\n\nOK restart: ok\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Request = httptest.NewRequest(http.MethodGet, "/restart?"+tt.query, nil)

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
