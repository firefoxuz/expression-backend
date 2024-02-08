package services

import "encoding/json"

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(Err bool, Message string, Data interface{}) *Response {
	return &Response{
		Error:   Err,
		Message: Message,
		Data:    Data,
	}
}

func (r *Response) ToJsonString() (string, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (r *Response) ToJsonBytes() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
