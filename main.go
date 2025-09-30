package main

import (
	
	"log"
	"net/http"
	"test/handler"	
)

func main() {
	http.HandleFunc("/cryptos", handler.GetCryptoDataHandler)

	log.Println("Server starting on port 8080...")
	log.Println("point d'accès API disponible sur http://localhost:8080/cryptos")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Impossible de démarrer le serveur: %s\n", err)
	}
}
