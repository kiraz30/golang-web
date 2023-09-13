package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (MyPage MyPage) SayHello(name string) string {
	return "Hello" + name + ", My Name is " + MyPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello " Bima"}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Setya",
	})

}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// global function len
func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Setya",
	})

}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// https://github.com/golang/go/blob/master/src/text/template/funcs.go

// Membuat global function sendiri untuk merubah string value menjadi Upper text
func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")

	//registrasi funcs
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{upper .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Setya",
	})

}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// function PIPELINES
func TemplateFunctionCreateGlobalPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")

	//registrasi funcs
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "hello" + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{sayHello .Name | upper }}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Setya",
	})

}

func TestTemplateFunctionCreateGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipeline(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
