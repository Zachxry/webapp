package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {

	mux := routes(&app)

	switch v := mux.(type) {
	case http.Handler:
		// do nothing; test passed

	default:
		t.Error(fmt.Sprintf("type is %T", v))
	}
}
