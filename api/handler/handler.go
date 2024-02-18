package handler

import (
	"encoding/json"
	"expression-backend/internal/models"
	"expression-backend/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Expression struct {
	Exp       string `json:"expression" validate:"required"`
	TimeLimit int    `json:"time_limit" validate:"required"`
}

func StoreExpression(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		r := services.NewResponse(true, "something went wrong", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
		return
	}

	exp := Expression{}

	err = json.Unmarshal(body, &exp)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		r := services.NewResponse(true, "something went wrong", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err = validate.Struct(exp); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		r := services.NewResponse(true, "error in validation", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
		return
	}

	expressionData := models.ExpressionData{
		Expression:   strings.ReplaceAll(exp.Exp, " ", ""),
		Result:       nil,
		IsProcessing: false,
		IsValid:      1,
		IsFinished:   false,
		TimeLimit:    exp.TimeLimit,
		CreatedAt:    time.Now().Format(time.DateTime),
		FinishedAt:   nil,
	}

	if err = expressionData.Store(); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		r := services.NewResponse(true, "something went wrong", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
		return
	}

	writer.WriteHeader(http.StatusOK)
	r := services.NewResponse(false, "expression is stored and soon will be calculated", nil)
	b, _ := r.ToJsonBytes()
	writer.Write(b)
}

func GetExpressions(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	expressions := models.ExpressionData{}
	data, err := expressions.GetAll()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		r := services.NewResponse(true, "something went wrong", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
		return
	}

	writer.WriteHeader(http.StatusOK)
	r := services.NewResponse(false, "", data)
	b, _ := r.ToJsonBytes()
	writer.Write(b)
}

func GetExpression(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	expressionId, ok := params["id"]

	if !ok {
		writer.WriteHeader(http.StatusInternalServerError)
		r := services.NewResponse(true, "something went wrong", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
		return
	}

	id, err := strconv.Atoi(expressionId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		r := services.NewResponse(true, "something went wrong", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
		return
	}

	expressions := models.ExpressionData{}
	data, err := expressions.FindById(id)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		r := services.NewResponse(true, "expression not found", nil)
		b, _ := r.ToJsonBytes()
		writer.Write(b)
		return
	}

	writer.WriteHeader(http.StatusOK)
	r := services.NewResponse(false, "", data)
	b, _ := r.ToJsonBytes()
	writer.Write(b)
}

func GetAgents(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	agents := services.GetAgents()

	var data []services.Agent = make([]services.Agent, 0)

	for id, t := range agents {
		data = append(data, services.Agent{
			Id:       id,
			LastPing: t.Add(time.Hour * 3).Format(time.DateTime),
			IsActive: t.Add(30 * time.Second).After(time.Now()),
		})
	}

	writer.WriteHeader(http.StatusOK)
	r := services.NewResponse(false, "", data)
	b, _ := r.ToJsonBytes()
	writer.Write(b)
}
