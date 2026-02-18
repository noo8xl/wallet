package main

import (
	"log"
	"net/http"

	"wallet/REST/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	walletHandler := handlers.NewWalletHandler()
	walletHandler.RegisterRoutes(mux)

	addr := "0.0.0.0:8080"
	log.Printf("REST wallet service listening on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start REST server: %v", err)
	}
}

