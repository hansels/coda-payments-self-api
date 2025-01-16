package main

import (
	"flag"
	"github.com/hansels/coda-payments-self-api/src/server"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	os.Exit(Main())
}

func Main() int {
	log.Println("Starting server...")

	// Get Port from Flag
	flag.Parse()
	port := flag.Arg(0)
	if port == "" {
		port = "3000"
	}

	delay := flag.Arg(1)
	if delay == "" {
		delay = "0"
	}

	delayInt, err := strconv.ParseInt(delay, 10, 64)
	if err != nil {
		log.Println("Invalid delay value")
		return 1
	}

	// Create new server
	api := server.New(&server.Opts{
		ListenAddress: ":" + port,
		Delay:         delayInt,
	})

	go api.Run()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-term:
		log.Println("Exiting gracefully...", s)
	}

	log.Println("👋")
	return 0
}
