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

func (h *Handler) InitRoutes(router *gin.Engine) {
	loginGroup := router.Group("")
	loginGroup.POST("/api/auth/google", h.VerifyGoogleToken)
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

func (h *Handler) fetchGoogleUserInfo(accessToken string) (UserInfo, error) {
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?access_token=%s", accessToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return UserInfo{}, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return UserInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, fmt.Errorf("failed to verify token")
	}

	var authResponse AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return UserInfo{}, err
	}

	// Log the complete user info for debugging
	logger.Debug("User info received from Google API: %+v", &authResponse)

	// Return the user's name in the desired JSON structure
	return UserInfo{Name: authResponse.Name}, nil
}
