package appholder

import (
	"backapper/app"
	"errors"
	"sync"
)

type AppHolder struct {
	m    sync.Mutex
	apps map[string]*app.App
}

func New(apps ...*app.App) *AppHolder {
	h := new(AppHolder)

	for _, a := range apps {
		h.apps[a.Name] = a
	}

	return h
}

func (h *AppHolder) getApp(name string) (*app.App, error) {
	h.m.Lock()
	defer h.m.Unlock()

	a, exists := h.apps[name]
	if !exists {
		return nil, errors.New("app doesn't exist")
	}

	return a, nil
}
