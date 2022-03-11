package main

import (
	"ginRouteAutoRegister/route"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	route.Route(router)
	srv := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//自定义退出时prestop逻辑，这里简单用等待5s
	time.Sleep(1 * time.Second)
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Println("Server Shutdown: ", err)
	//}
	//select {
	//case <-ctx.Done():
	//	log.Println("timeout of 5 seconds.")
	//}
	log.Println("Server exiting")
}
