package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// The program has been exited, but no clearnup work
	var closing = make(chan struct{})

	// Clearnup work is done
	var closed = make(chan struct{})

	// Mock some business work
	go func() {
		for {
			select {
			case <-closing:
				fmt.Println("Stop work!")
				return
			default:
				// Business calculate...
				fmt.Println("I'am working well php is very good...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Process wait CTRL+C interrupt signal
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	fmt.Println("Received CTRL+C inetrrupt signal")

	// Stop work
	close(closing)

	// Executing before exit cleanup
	go doCleanup(closed)

	select {
	case <-closed:
		fmt.Println("Cleanup work is done!")
	case <-time.After(time.Second * 5):
		fmt.Println("Cleanup timeout, cancle wait")
	}

	fmt.Println("Graceful shutdown!")

}

// Do cleanup
func doCleanup(closed chan struct{}) {
	time.Sleep(time.Minute)
	close(closed)
}
