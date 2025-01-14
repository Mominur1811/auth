package utils

import "net/http"

func SendData(w http.ResponseWriter, statusCode int, data interface{}) {
	SendJson(w, statusCode, data)
}
