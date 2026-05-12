package main

import (
	"log"

	"github.com/Dukvaha27/flash-score/gateway/internal/transport"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	transport.RegisterProxies(router)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
