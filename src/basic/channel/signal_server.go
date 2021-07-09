package main

import (
	"fmt"
	"net/http"
)

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})

	go func() {
		done <- serverDebug(stop)
	}()

	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Serve{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop // wait for stop signal
		s.Shutdown(context.Backgroud())
	}()
	return s.ListenAndServe()
}
