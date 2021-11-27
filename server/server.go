package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()
	Router(router)
	serve := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
