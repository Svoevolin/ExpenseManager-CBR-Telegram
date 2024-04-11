package cbr_gateway

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Svoevolin/workshop_1_bot/internal/domain"
	"golang.org/x/text/encoding/charmap"
)

type Gateway struct {
	client *http.Client
}

func New() *Gateway {
	return &Gateway{
		client: http.DefaultClient,
	}
}

func (gate *Gateway) FetchRates(ctx context.Context, date time.Time) ([]domain.Rate, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	log.Printf("start receiving exchange rates on %s", date.Format("02/01/2006"))

	url := fmt.Sprintf("https://www.cbr.ru/scripts/XML_daily.asp?date_req=%s", date.Format("02/01/2006"))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	resp, err := gate.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get a list of currencies on the date %s", date.Format("02/01/2006"))
	}

	defer resp.Body.Close()

	d := xml.NewDecoder(resp.Body)
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("Unknown charset: %s", charset)
		}
	}

	var cbrRates Rates
	if err = d.Decode(&cbrRates); err != nil {
		return nil, err
	}

	rates := make([]domain.Rate, len(cbrRates.Currencies))
	for _, rate := range cbrRates.Currencies {
		rates = append(rates, domain.Rate{
			Code:     rate.CharCode,
			Original: strings.Replace(rate.Value, ",", ".", 1),
			Nominal:  rate.Nominal,
		})
	}
	return rates, nil
}
