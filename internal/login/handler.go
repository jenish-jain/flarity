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

	// Respond with the user's name in the desired JSON structure
	ctx.JSON(http.StatusOK, userInfo)
}

func (h *Handler) fetchGoogleUserInfo(accessToken string) (map[string]string, error) {
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

	// Log the complete user info for debugging
	logger.Debug("User info received from Google API: %+v", userInfo)

	// Extract the user's name from the response
	name, ok := userInfo["name"].(string)
	if !ok {
		return nil, fmt.Errorf("user's name not found in response")
	}

	// Return the user's name in the desired JSON structure
	return map[string]string{"name": name}, nil
}
