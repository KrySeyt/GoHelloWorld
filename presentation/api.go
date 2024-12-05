package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"packname/app"
	"strings"
)

func ensure_method_allowed(methodType string, writer http.ResponseWriter, request *http.Request) bool {
	if strings.ToUpper(request.Method) != strings.ToUpper(methodType) {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(writer, "Method not allowed")
		return false
	}
	return true
}

type HelloInput struct {
	Name string
}

type HelloHttpApi struct {
	sayHello *app.SayHello
	getHello *app.GetHello
}

func (self *HelloHttpApi) GetHello(writer http.ResponseWriter, request *http.Request) {
	if !ensure_method_allowed("GET", writer, request) {
		return
	}

	name := request.URL.Query().Get("name")

	hello := self.getHello.Execute(name)

	fmt.Fprint(writer, hello)
}

func (self *HelloHttpApi) SayHello(writer http.ResponseWriter, request *http.Request) {
	if !ensure_method_allowed("POST", writer, request) {
		return
	}

	var input HelloInput
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&input)

	self.sayHello.Execute(input.Name)

	fmt.Fprint(writer, "success")
}

func CreateHelloHttpApi(
	sayHello *app.SayHello,
	getHello *app.GetHello,
	) *HelloHttpApi {
	return &HelloHttpApi{
		sayHello: sayHello,
		getHello: getHello,
	}
}
