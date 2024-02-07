package handlers

import (
	"encoding/json"
	"expression-backend/internal/app/services"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Expression struct {
	Exp string `json:"expression" validate:"required"`
}

func StoreExpression(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("something went wrong"))
	}

	exp := Expression{}

	json.Unmarshal(body, &exp)

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(exp)

	fmt.Println(err)
	fmt.Println(exp.Exp)
	fmt.Println(services.IsValidMathExpression(exp.Exp))
}

func GetExpressions(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("method get"))
}

func GetExpression(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	expressionId, ok := params["id"]

	if !ok {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("something went wrong"))
	}

	writer.Write([]byte(expressionId))
}
