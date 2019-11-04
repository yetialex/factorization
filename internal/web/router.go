package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func newRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/factorize/{number}", FactorizeHandler).Methods("GET")
	r.HandleFunc("/factorize", FactorizeHandler).Methods("POST")
	return r
}

func Start() *http.Server {
	r := newRouter()
	listenAddr := ":8080"
	envPort := os.Getenv("listen_port")
	if envPort != "" {
		listenAddr = fmt.Sprintf(":%s", envPort)
	}
	log.Println("listening on", listenAddr)
	srv := &http.Server{
		Addr:         listenAddr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	return srv
}

func GracefulShutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Println("graceful shutdown error: ", err)
	}
}
