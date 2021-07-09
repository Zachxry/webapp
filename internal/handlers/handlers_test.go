package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
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
	{"order-summary", "/order-summary", "GET", []postData{}, http.StatusOK},
	{"confirm", "/confirm", "GET", []postData{}, http.StatusOK},
	{"post-confirm", "/confirm", "POST", []postData{
		{key: "first_name", value: "Zach"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "zach@example.com"},
		{key: "phone", value: "210-222-2222"},
	}, http.StatusOK},
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
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expeced %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
