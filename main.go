package main

import (
	"log"
	"flag"
	"github.com/AleksanderWWW/email-domain-verification/api"
	"github.com/gin-gonic/gin"
)

func main() {
	listenAddr := flag.String("listenaddr", ":5555", "api port to listen on")
	flag.Parse()

	router := gin.Default()

	router.POST("api/v1/verify", api.HandleVerification)

	log.Fatal(router.Run(*listenAddr))
}
