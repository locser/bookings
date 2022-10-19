package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/locser/bookings/internal/config"
)

//go get github.com/bmizerany/pat to download

func TestRoutes(t *testing.T) {

	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// TODO:
	default:
		t.Errorf("type is not *chi.Mux is %T", v)
	}
}
