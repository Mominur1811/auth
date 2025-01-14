package utils

import (
	"auth-repo/logger"
	"encoding/json"
	"log/slog"
	"net/http"
)

func SendJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	str, err := json.Marshal(data)
	if err != nil {
		slog.Error(err.Error(), logger.Extra(map[string]any{
			"data": data,
		}))
	}

	w.WriteHeader(status)
	w.Write(str)
}
