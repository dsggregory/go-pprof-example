package main

import (
	"go-pprof-example/pkg/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigdone := make(chan os.Signal)

	signal.Notify(sigdone, syscall.SIGINT)

	router := api.NewServer(":8080", true)

	go func() {
		if err := router.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<- sigdone
}
