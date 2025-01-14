package handlers

import (
	"auth-repo/rest/utils"
	"net/http"
)

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, 200, map[string]interface{}{
		"hello": "world",
	})
}
