package http

import (
	"arvancloud-challenge-app/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	requestUC usecase.RequestUseCase
}

func NewRequestHandler(r *gin.Engine, requestUC usecase.RequestUseCase) {
	handler := &RequestHandler{requestUC: requestUC}

	r.GET("/log", handler.LogRequest)
}

func (h *RequestHandler) LogRequest(c *gin.Context) {
	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()

	err := h.requestUC.LogRequest(ip, userAgent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "logged"})
}
