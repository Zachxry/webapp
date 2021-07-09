package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/url", nil)
	form := New(r.PostForm)

	isValid := form.Valid() // testing the Valid function which should return true here.
	if !isValid {
		t.Error("got invalid when form should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/url", nil)

	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/url", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/url", nil)
	form := New(r.PostForm)

	has := form.Has("someDataHere")
	if !has { // incorrect
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)
	has = form.Has("a")
	if has { // incorrect
		t.Error("shows form does not have field when it should")
	}

}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/url", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)

	form.MinLength("some_length", 100)
	if form.Valid() {
		t.Error("shows min length of 100 met when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)

	form.MinLength("another_field", 3)
	if !form.Valid() {
		t.Error("shows min length of 3 is not met when it is")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "test@test.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when it was vaild")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "testingtest.com")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got a valid email when it was not vaild")
	}

}
