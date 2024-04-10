package cbr_gateway

import "net/http"

type Gateway struct {
	client *http.Client
}
