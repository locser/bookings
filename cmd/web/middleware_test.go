package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myHan myHandler

	h := NoSurf(&myHan)

	switch v := h.(type) {
	case http.Handler:
		// TODO:

	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}

}

func TestSessionLoad(t *testing.T) {
	var myHan myHandler

	h := NoSurf(&myHan)

	switch v := h.(type) {
	case http.Handler:
		// TODO:

	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}

}
