package appservice

import (
	"backapper/app"
	"backapper/app/appholder"
	"errors"
	"log"
	"os/exec"
	"testing"
)

func TestAppService_Restart(t *testing.T) {
	logger := log.Default()
	service := New(appholder.New(
		&app.App{Name: "ok", Restart: "echo restarted"},
		&app.App{Name: "fail", Restart: "echo failed >&2; exit 1"},
	), logger)

	tests := []struct {
		name          string
		appName       string
		wantOutput    string
		wantErr       bool
		wantExitError bool
	}{
		{
			name:       "success",
			appName:    "ok",
			wantOutput: "restarted\n",
		},
		{
			name:          "script exit non-zero",
			appName:       "fail",
			wantErr:       true,
			wantExitError: true,
		},
		{
			name:    "app not found",
			appName: "missing",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := service.Restart(tt.appName)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				var exitErr *exec.ExitError
				if tt.wantExitError != errors.As(err, &exitErr) {
					t.Fatalf("errors.As(*exec.ExitError) = %v, want %v", errors.As(err, &exitErr), tt.wantExitError)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if output != tt.wantOutput {
				t.Fatalf("output = %q, want %q", output, tt.wantOutput)
			}
		})
	}
}

func TestAppService_HealthCheck(t *testing.T) {
	logger := log.Default()
	service := New(appholder.New(
		&app.App{Name: "ok", HealthCheck: "echo healthy"},
		&app.App{Name: "fail", HealthCheck: "exit 1"},
		&app.App{Name: "no-config"},
	), logger)

	tests := []struct {
		name          string
		appName       string
		wantOutput    string
		wantErr       bool
		wantExitError bool
	}{
		{
			name:       "success",
			appName:    "ok",
			wantOutput: "healthy\n",
		},
		{
			name:          "script exit non-zero",
			appName:       "fail",
			wantErr:       true,
			wantExitError: true,
		},
		{
			name:    "app not found",
			appName: "missing",
			wantErr: true,
		},
		{
			name:    "healthcheck not configured",
			appName: "no-config",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := service.HealthCheck(tt.appName)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				var exitErr *exec.ExitError
				if tt.wantExitError != errors.As(err, &exitErr) {
					t.Fatalf("errors.As(*exec.ExitError) = %v, want %v", errors.As(err, &exitErr), tt.wantExitError)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if output != tt.wantOutput {
				t.Fatalf("output = %q, want %q", output, tt.wantOutput)
			}
		})
	}
}
