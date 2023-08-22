package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello my name %s", name)
	}
}
func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	//cek response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)
	fmt.Println(stringBody)
}

// Multi query paramater
func SayMultiHello(w http.ResponseWriter, r *http.Request) {
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")
	fmt.Fprintf(w, "Nama Saya %s %s", first_name, last_name)
}

func TestMultiquery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Bima&last_name=NUgraha", nil)
	recorder := httptest.NewRecorder()

	SayMultiHello(recorder, request)
	//cek response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)
	fmt.Println(stringBody)
}

// Multy value Query Paramater
func MultiValueQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]

	fmt.Fprintf(w, strings.Join(names, " "))
}

func TestMultiValueQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Bima&name=Setya&name=Nugraha", nil)
	recorder := httptest.NewRecorder()

	MultiValueQuery(recorder, request)
	//cek response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)
	fmt.Println(stringBody)
}
