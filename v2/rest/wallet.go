package rest

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"path"
)

// WalletService manages data flow for the Wallet API endpoint
type WalletService struct {
	requestFactory
	Synchronous
}

// All returns all orders for the authenticated account.
func (s *WalletService) Wallet() (*bitfinex.WalletSnapshot, error) {
	raw, err := s.requestFactory.MakeNewAuthenticatedRequest(path.Join("wallets"), s)
	if err != nil {
		return nil, err
	}

	os, err := bitfinex.NewWalletSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}

	return os, nil
}
