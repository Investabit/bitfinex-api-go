package rest

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"net/http"
)

// PositionService manages the Position endpoint.
type PositionService struct {
	requestFactory
	Synchronous
}

// All returns all positions for the authenticated account.
func (s *PositionService) All() (os *bitfinex.PositionSnapshot, resp *http.Response, err error) {
	var raw []interface{}
	raw, resp, err = s.requestFactory.MakeNewAuthenticatedRequest("positions", s)

	if err != nil {
		return
	}

	os, err = bitfinex.NewPositionSnapshotFromRaw(raw)
	if err != nil {
		return
	}

	return
}
