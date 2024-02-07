package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func StoreExpression(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("method post"))
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
