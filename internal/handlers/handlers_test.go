package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"coffee", "/coffee", "GET", []postData{}, http.StatusOK},
	{"cassava-cake", "/cassava-cake", "GET", []postData{}, http.StatusOK},
	{"order", "/order", "GET", []postData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expeced %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			fmt.Println("placeholder for test")
		}
	}
}
