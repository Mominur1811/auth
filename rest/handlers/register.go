package handlers

import (
	"auth-repo/rest/utils"
	"auth-repo/types"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	var userInfo types.UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		slog.Error("Failed to decode user info", map[string]interface{}{
			"userInfo": userInfo,
			"err":      err.Error(),
		})
		utils.SendError(w, http.StatusBadRequest, "Failed to register", "")
		return
	}

	if err := utils.Validate(userInfo); err != nil {
		slog.Error("Failed to validate user info", map[string]interface{}{
			"userInfo": userInfo,
			"err":      err.Error(),
		})
		utils.SendError(w, http.StatusBadRequest, "Failed to register", "")
		return
	}

	fmt.Println(userInfo)
	user, err := h.service.Register(r.Context(), userInfo)
	if err != nil {
		slog.Error("Failed to register user", map[string]interface{}{
			"userInfo": userInfo,
			"err":      err.Error(),
		})
		utils.SendError(w, http.StatusInternalServerError, "Failed to register", "")
		return
	}

	utils.SendData(w, http.StatusOK, user)
}
