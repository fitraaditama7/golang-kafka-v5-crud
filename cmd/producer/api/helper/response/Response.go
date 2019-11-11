package response

import (
	"database/sql"
	"encoding/json"
	"golang-kafka-v5-crud/cmd/producer/api/libs"
	"net/http"
)

type Responses struct {
	Error   bool        `json:"error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(w http.ResponseWriter, code int, msg string, payload interface{}) {
	var result Responses

	if code >= 400 {
		result.Error = true
		result.Code = code
		result.Message = "Error"
		result.Data = payload
	} else {
		result.Error = false
		result.Code = code
		if msg == "" {
			msg = "Success"
		}
		result.Message = msg
		result.Data = payload
	}

	response, _ := json.Marshal(result)
	libs.KafkaProducer(string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ErrorCustomStatus(w http.ResponseWriter, code int, err error) {
	if err == sql.ErrNoRows {
		Response(w, http.StatusNotFound, "Error", err.Error())
		return
	}

	Response(w, code, "Error", err.Error())
	return
}
