package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// headaer informasi tambahan yang bisa dikirim oleh client ke server atau sebaliknya

// menangkap data header
func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprintf(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("content-type", "aplication/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// menambahkan data header
func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Bima Setya Nugraha")
	fmt.Fprintf(w, "Ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("content-type", "aplication/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)
	// response := recorder.Result()
	// body, _ := io.ReadAll(response.Body)
	fmt.Println(recorder.Header().Get("X-POWERED-BY"))
}
