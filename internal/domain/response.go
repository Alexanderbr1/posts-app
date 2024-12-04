package domain

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Print(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}
