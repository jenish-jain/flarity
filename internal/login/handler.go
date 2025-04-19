package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jenish-jain/logger"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	router.POST("/api/auth/google", h.VerifyGoogleToken)
}

func (h *Handler) VerifyGoogleToken(ctx *gin.Context) {
	var requestBody struct {
		Token string `json:"token"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		logger.ErrorWithCtx(ctx, "Invalid request body", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	userInfo, err := h.fetchGoogleUserInfo(requestBody.Token)
	if err != nil {
		logger.ErrorWithCtx(ctx, "Failed to verify token", "err", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	// Handle user login or registration logic here
	ctx.JSON(http.StatusOK, gin.H{"message": "Authentication successful", "user": userInfo})
}

func (h *Handler) fetchGoogleUserInfo(accessToken string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?access_token=%s", accessToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to verify token")
	}

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
