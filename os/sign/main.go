package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var firstSigusr1 = true

func main() {
	// 忽略 Control-C (SIGINT)
	// os.Interrupt 和 syscall.SIGINT 是同义词
	signal.Ignore(os.Interrupt)

	c1 := make(chan os.Signal, 2)
	// Notify SIGHUP
	signal.Notify(c1, syscall.SIGHUP)
	// Notify SIGUSR1
	signal.Notify(c1, syscall.SIGUSR1)
	go func() {
		for {
			switch <-c1 {
			case syscall.SIGHUP:
				fmt.Println("sighup, reset sighup")
				signal.Reset(syscall.SIGHUP)
			case syscall.SIGUSR1:
				if firstSigusr1 {
					fmt.Println("first usr1, notify interrupt which had ignore!")
					c2 := make(chan os.Signal, 1)
					// Notify Interrupt
					signal.Notify(c2, os.Interrupt)
					go handlerInterrupt(c2)
				}
			}
		}
	}()

	select {}
}

func handlerInterrupt(c <-chan os.Signal) {
	for {
		switch <-c {
		case os.Interrupt:
			fmt.Println("signal interrupt")
		}
	}
}
