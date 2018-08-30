package bitfinex

import (
	"net/http"
)

type BalancesService struct {
	client *Client
}

type WalletBalance struct {
	Type      string
	Currency  string
	Amount    string
	Available string
}

// GET balances
func (b *BalancesService) All() ([]WalletBalance, *http.Response, error) {
	balances := make([]WalletBalance, 3)
	resp, err := b.client.authenticatedAndDoRequest("GET", "balances", nil, &balances)
	if err != nil {
		if resp != nil {
			return nil, resp.Response, err
		}
		return nil, nil, err
	}

	return balances, resp.Response, nil
}
