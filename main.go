package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// hooking to ctrl-c os signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nGood bye!")
		// dont forget to exit, very important!
		os.Exit(0)
	}()

	for {
		fmt.Println("Hello there!")
		time.Sleep(time.Second * 1)
	}
}
