package models

import (
	"expression-backend/internal/database"
)

type (
	booleanOrNil = interface{}
	stringOrNil  = interface{}
	intOrNil     = interface{}
)

type ExpressionData struct {
	Id           int          `json:"id" db:"id"`
	Expression   string       `json:"expression" db:"expression"`
	Result       intOrNil     `json:"result" db:"result"`
	IsProcessing bool         `json:"is_processing" db:"is_processing"`
	IsTimeLimit  booleanOrNil `json:"is_time_limit" db:"is_time_limit"`
	IsValid      booleanOrNil `json:"is_valid" db:"is_valid"`
	IsFinished   bool         `json:"is_finished" db:"is_finished"`
	TimeLimit    int          `json:"time_limit" db:"time_limit"`
	CreatedAt    string       `json:"created_at" db:"created_at"`
	FinishedAt   stringOrNil  `json:"finished_at" db:"finished_at"`
}

func (d *ExpressionData) Store() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	_, err = db.NamedExec(`INSERT INTO expressions (expression, result, is_processing, is_time_limit, is_valid, is_finished, time_limit, created_at, finished_at)
        VALUES (:expression, :result, :is_processing, :is_time_limit, :is_valid, :is_finished, :time_limit, :created_at, :finished_at)`, d)

	return err
}

func (d *ExpressionData) GetAll() (*[]ExpressionData, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var data []ExpressionData
	err = db.Select(&data, "SELECT * FROM expressions ORDER BY id desc")

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *ExpressionData) FindById(id int) (*ExpressionData, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var data ExpressionData
	err = db.Get(&data, "SELECT * FROM expressions where id=$1 limit 1", id)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *ExpressionData) Update() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}
	_, err = db.NamedExec("UPDATE expressions set finished_at=current_timestamp, is_finished=true, is_time_limit=:is_time_limit, is_valid=:is_valid, result=:result  where id=:id", d)
	return err
}

func (d *ExpressionData) GetNotFinished() (*[]ExpressionData, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var data []ExpressionData
	err = db.Select(&data, "SELECT * FROM expressions where is_finished = false")

	if err != nil {
		return nil, err
	}

	return &data, nil
}
