package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/models"
)

const (
	coinGeckoAPIURL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1"
)

func FetchCryptos() ([]models.Crypto, error) {
	resp, err := http.Get(coinGeckoAPIURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from CoinGecko: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var cryptos []models.Crypto
	if err := json.NewDecoder(resp.Body).Decode(&cryptos); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	
	return cryptos, nil
}