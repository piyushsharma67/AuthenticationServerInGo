package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status    bool    `json:"status"`
	ErrorMessage  string `json:"message"`
}

func CreateErrorResponse(err error, statusCode int, w http.ResponseWriter) {
	resp := &ErrorResponse{false, err.Error()}
	respJSON, _ := json.Marshal(resp)
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.WriteHeader(statusCode)
	w.Write(respJSON)
}

type SuccessResponse struct {
	Status bool         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func CreateSuccessResponse(message string, data interface{}, statusCode int, w http.ResponseWriter) {
	resp := &SuccessResponse{true, message, data}
	respJSON, _ := json.Marshal(resp)
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.WriteHeader(statusCode)
	w.Write((respJSON))
}

func CreateLoginJwtToken(){
	
}