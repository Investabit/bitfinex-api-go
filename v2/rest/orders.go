package rest

import (
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"net/http"
	"path"
	"strconv"
)

// OrderService manages data flow for the Order API endpoint
type OrderService struct {
	requestFactory
	Synchronous
}

// All returns all orders for the authenticated account.
func (s *OrderService) All(symbol string) (*bitfinex.OrderSnapshot, *http.Response, error) {
	raw, resp, err := s.requestFactory.MakeNewAuthenticatedRequest(path.Join("orders", symbol), s)
	if err != nil {
		return nil, resp, err
	}

	os, err := bitfinex.NewOrderSnapshotFromRaw(raw)
	if err != nil {
		return nil, resp, err
	}

	return os, resp, nil
}

// Status retrieves the given order from the API. This is just a wrapper around
// the All() method, since the API does not provide lookup for a single Order.
func (s *OrderService) Status(orderID int64) (o *bitfinex.Order, resp *http.Response, err error) {
	os, resp, err := s.All("")

	if err != nil {
		return o, resp, err
	}

	if len(os.Snapshot) == 0 {
		return o, resp, bitfinex.ErrNotFound
	}

	for _, e := range os.Snapshot {
		if e.ID == orderID {
			return e, resp, nil
		}
	}

	return o, resp, bitfinex.ErrNotFound
}

// All returns all orders for the authenticated account.
func (s *OrderService) History(symbol string) (*bitfinex.OrderSnapshot, *http.Response, error) {
	if symbol == "" {
		return nil, nil, fmt.Errorf("symbol cannot be empty")
	}
	raw, resp, err := s.requestFactory.MakeNewAuthenticatedRequest(path.Join("orders", symbol, "hist"), s)
	if err != nil {
		return nil, resp, err
	}

	os, err := bitfinex.NewOrderSnapshotFromRaw(raw)
	if err != nil {
		return nil, resp, err
	}

	return os, resp, nil
}

// All returns all trades for a given order for the authenticated account.
func (s *OrderService) OrderTrades(symbol string, orderID int64) (*bitfinex.OrderTradeSnapshot, *http.Response, error) {
	if symbol == "" {
		return nil, nil, fmt.Errorf("symbol cannot be empty")
	}

	id := strconv.FormatInt(orderID, 10)

	raw, resp, err := s.requestFactory.MakeNewAuthenticatedRequest(path.Join("order", symbol+":"+id, "trades"), s)
	if err != nil {
		return nil, resp, err
	}

	os, err := bitfinex.NewOrderTradeSnapshotFromRaw(raw)
	if err != nil {
		return nil, resp, err
	}

	return os, resp, nil
}
