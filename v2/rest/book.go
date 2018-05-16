package rest

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type BookService struct {
	Synchronous
}

func (b *BookService) All(symbol string, precision bitfinex.BookPrecision, priceLevels int) (*bitfinex.BookUpdateSnapshot, *http.Response, error) {
	req := NewRequestWithMethod(path.Join("book", symbol, string(precision)), "GET")
	req.Params = make(url.Values)
	req.Params.Add("len", strconv.Itoa(priceLevels))
	raw, resp, err := b.Request(req)

	if err != nil {
		return nil, resp, err
	}

	data := make([][]float64, 0, len(raw))
	for _, ifacearr := range raw {
		if arr, ok := ifacearr.([]interface{}); ok {
			sub := make([]float64, 0, len(arr))
			for _, iface := range arr {
				if flt, ok := iface.(float64); ok {
					sub = append(sub, flt)
				}
			}
			data = append(data, sub)
		}
	}

	book, err := bitfinex.NewBookUpdateSnapshotFromRaw(symbol, string(precision), data)
	if err != nil {
		return nil, resp, err
	}

	return book, resp, nil
}
