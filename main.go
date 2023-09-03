package main

import (
	"github.com/AleksanderWWW/email-domain-verification/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("api/v1/verify", api.HandleVerification)

	router.Run(":8080")
}
