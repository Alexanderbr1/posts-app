package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"posts-app/internal/domain"
	"strings"
)

func (h *Handler) siqnUp(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.Authorization.CreateUser(c, input)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) siqnIn(c *gin.Context) {
	var input domain.SignInInput
	if err := c.BindJSON(&input); err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	accessToken, refreshToken, err := h.service.Authorization.SignIn(c, input)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": accessToken,
	})
}

func (h *Handler) refresh(c *gin.Context) {
	cookie, err := c.Cookie("refresh-token")
	if err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token := strings.ReplaceAll(cookie, "'", "")
	accessToken, refreshToken, err := h.service.RefreshTokens(c, token)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": accessToken,
	})
}
