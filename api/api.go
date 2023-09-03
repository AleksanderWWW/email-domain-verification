package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerificationRequest struct {
	Domain string `json:"domain"`
}


func HandleVerification(c *gin.Context) {
	var vr VerificationRequest

	if err := c.BindJSON(&vr); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	domainInfo, err := CheckDomain(vr.Domain)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, &domainInfo)
}
