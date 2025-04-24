package app

import (
	"context"
	"fullstacktest/db"
	"fullstacktest/handler"
	"fullstacktest/middleware"
	"fullstacktest/repository"
	"fullstacktest/usecase"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("error connect DB: %s", err)
	}
	r := gin.New()
	r.ContextWithFallback = true //request cancellation

	r.Use(gin.Recovery())

	r.Use(middleware.RequestIDMiddleware())

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	cr := repository.NewClientRepo(db)
	cuc := usecase.NewClientUsecase(cr)
	ch := handler.NewClientHandler(cuc)

	r.GET("/myclients", ch.GetClients)
	r.POST("/myclients", ch.InsertClient)
	r.PATCH("/myclients", ch.UpdateClient)
	r.DELETE("/myclients/:slug", ch.DeleteClient)

	//Graceful shutdown
	srv := &http.Server{
		Addr:    os.Getenv("APP_PORT"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
