package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := setup(); err != nil {
		panic(err)
	}
}

func setup() error {
	conn, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}

	defer conn.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello From Kube\n"))
		w.Write([]byte(fmt.Sprintf("Request IP: %s\n", r.RemoteAddr)))
		w.Write([]byte("Easy kube app\n"))
	})

	s := http.Server{
		Handler: mux,
	}

	gS := make(chan os.Signal, 1)
	signal.Notify(gS, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Println("Starting web server...")
		if err := s.Serve(conn); err != nil && err != http.ErrServerClosed {
			log.Println("Serv err: ", err)
		}
	}()

	<-gS

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	err = s.Shutdown(ctx)
	log.Println("Done!")

	return err
}
