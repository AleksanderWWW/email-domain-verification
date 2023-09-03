package main

import "github.com/gin-gonic/gin"
import "github.com/AleksanderWWW/email-domain-verification/api"


func main() {
	router := gin.Default()

  	router.POST("api/v1/verify", api.HandleVerification)

	router.Run(":8080")
}
