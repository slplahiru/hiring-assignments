package main

import (
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServeRandomFileBasic(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	serveRandomFile(w, req)
	res := w.Result()

	defer res.Body.Close()

	typ := res.Header.Get("Content-Type")
	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if typ != "image/png" && typ != "application/pdf" {
		t.Errorf("Expected png or pdf but got '%v'", typ)
	}
}

func TestServeRandomFileNumbered(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	req := httptest.NewRequest("GET", "/987123", nil)
	w := httptest.NewRecorder()

	serveRandomFile(w, req)
	res := w.Result()

	defer res.Body.Close()

	typ := res.Header.Get("Content-Type")
	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if typ != "image/png" && typ != "application/pdf" {
		t.Errorf("Expected png or pdf but got '%v'", typ)
	}
}

func TestServeRandomFileLettered(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	expected_err := "Need an integer\n"
	expected_err_code := 400

	req := httptest.NewRequest("GET", "/asdafs", nil)
	w := httptest.NewRecorder()

	serveRandomFile(w, req)
	res := w.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	err_code := res.StatusCode
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected_err {
		t.Errorf("Expected error '%v' but got '%v'", expected_err, string(data))
	}
	if err_code != expected_err_code {
		t.Errorf("Expected error '%v' but got '%v'", expected_err_code, err_code)
	}
}

func TestServeRandomFileMixed(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	expected_err := "Need an integer\n"
	expected_err_code := 400

	req := httptest.NewRequest("GET", "/asf1234", nil)
	w := httptest.NewRecorder()

	serveRandomFile(w, req)
	res := w.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	err_code := res.StatusCode
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected_err {
		t.Errorf("Expected error '%v' but got '%v'", expected_err, string(data))
	}
	if err_code != expected_err_code {
		t.Errorf("Expected error '%v' but got '%v'", expected_err_code, err_code)
	}
}

func TestServeRandomFilePOST(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	expected_err := "Not supported\n"
	expected_err_code := 403

	req := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()

	serveRandomFile(w, req)
	res := w.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	err_code := res.StatusCode
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected_err {
		t.Errorf("Expected error '%v' but got '%v'", expected_err, string(data))
	}
	if err_code != expected_err_code {
		t.Errorf("Expected error '%v' but got '%v'", expected_err_code, err_code)
	}
}

func TestHealthCheck(t *testing.T) {
	expected_msg := "Health OK!"
	expected_err_code := 200

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	healthCheck(w, req)
	res := w.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	err_code := res.StatusCode
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected_msg {
		t.Errorf("Expected error '%v' but got '%v'", expected_msg, string(data))
	}
	if err_code != expected_err_code {
		t.Errorf("Expected error '%v' but got '%v'", expected_err_code, err_code)
	}
}
