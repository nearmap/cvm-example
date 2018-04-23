package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const html = `
			<!DOCTYPE html>
			<html>
			<body style="background-position: center;background-repeat: no-repeat;background-image: url('https://1.bp.blogspot.com/-wwm4xAAmbC0/VzY4flps3oI/AAAAAAAAGhQ/0agEuieNpJAClu4Vb0MeaOm5I5UY4r-2QCLcB/s1600/macaco.gif');">
			<h1>Echo: %s!</h1>
			</body>
			</html>
			`

func main() {

	var portVar int
	flag.IntVar(&portVar, "port", 80, "Port number for service")
	flag.Parse()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	mux := http.NewServeMux()
	mux.HandleFunc("/echo", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf(html, req.URL.Query().Get("msg")))
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", portVar),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 1 * time.Minute,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("Err:%v", err)
		}
	}()

	<-stop
	fmt.Println("Shutting down the server...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	fmt.Println("Server gracefully stopped")

}
