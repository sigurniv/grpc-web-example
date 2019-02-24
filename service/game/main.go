package main

import (
	"log"
	"fmt"
	"net"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"github.com/sigurniv/grpc-web-golang/service/game/server"
	"github.com/sigurniv/grpc-web-golang/service/game/gamepb"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"net/http"
	"time"
	"strings"
	"github.com/gorilla/websocket"
	"path"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	srv := server.New()

	fmt.Println("Game server Started")

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	gamepb.RegisterGameServiceServer(s, srv)

	go func() {
		log.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	go func() {
		log.Println("Starting grpc-web server")
		wrappedGrpc := grpcweb.WrapServer(s)

		httpSrv := &http.Server{
			// These interfere with websocket streams, disable for now
			// ReadTimeout: 5 * time.Second,
			// WriteTimeout: 10 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			Addr:              ":50053",
			IdleTimeout:       120 * time.Second,
			Handler:

			grpcTrafficSplitter(
				//folderReader(
				http.FileServer(http.Dir("static/")).ServeHTTP,
				//),
				wrappedGrpc,
			),
		}

		go func() {
			log.Fatal(httpSrv.ListenAndServe())
		}()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	log.Println("Stopping the server")
	s.Stop()
	log.Println("Closing the listener")
	lis.Close()
	log.Println("End of Program")
}

func folderReader(fn http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			// Use contents of index.html for directory, if present.
			r.URL.Path = path.Join(r.URL.Path, "index.html")
		}
		fn(w, r)
	})
}

func grpcTrafficSplitter(fallback http.HandlerFunc, grpcHandler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Redirect gRPC and gRPC-Web requests to the gRPC Server
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") ||
			websocket.IsWebSocketUpgrade(r) {
			if ct := r.Header.Get("Accept"); ct != "" {
				w.Header().Set("Content-Type", ct)
			}

			grpcHandler.ServeHTTP(w, r)
		} else {
			fallback(w, r)
		}
	})
}
