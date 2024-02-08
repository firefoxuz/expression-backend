package handler

import (
	"encoding/json"
	"expression-backend/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Expression struct {
	Exp string `json:"expression" validate:"required"`
}

func StoreExpression(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		r := services.NewResponse(true, "something went wrong", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
	}

	exp := Expression{}

	err = json.Unmarshal(body, &exp)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		r := services.NewResponse(true, "something went wrong", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err = validate.Struct(exp); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		r := services.NewResponse(true, "expression is not presented", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
	}

	if !services.IsValidExpression(exp.Exp) {
		writer.WriteHeader(http.StatusBadRequest)
		r := services.NewResponse(true, "expression is not valid", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
	}

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
