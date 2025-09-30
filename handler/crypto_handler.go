package handler


import (
	"encoding/json"
	"net/http"
	"test/services"
)

func GetCryptoDataHandler(w http.ResponseWriter, r *http.Request) {
	cryptos, err := services.FetchCryptos()
	if err != nil {
		http.Error(w, "Failed to fetch cryptocurrencies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cryptos); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}