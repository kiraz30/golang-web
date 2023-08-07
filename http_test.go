package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hi", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	//cek response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	//convert body ke string
	stringBody := string(body)

	fmt.Println(stringBody)
}
