package main

import (
	"os"
	"os/signal"

	"github.com/yetialex/factorization/internal/web"
)

func main() {
	srv := web.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	web.GracefulShutdown(srv)
}
