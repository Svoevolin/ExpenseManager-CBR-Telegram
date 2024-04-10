package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringAmountToKopecks(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		kopecks int64
		amount  string
		err     error
	}{
		{
			name:    "ok",
			input:   "123.45",
			kopecks: 12345,
			amount:  "123.45",
		},
		{
			name:    "another symbols",
			input:   "123,456,789.10",
			kopecks: 12345678910,
			amount:  "123456789.10",
		},
		{
			name:    "more than 2 digits after dot",
			input:   "123.45678",
			kopecks: 12345,
			amount:  "123.45",
		},
		{
			name:    "less than 2 digits after dot",
			input:   "123.4",
			kopecks: 12340,
			amount:  "123.40",
		},
		{
			name:    "no dot",
			input:   "123",
			kopecks: 12300,
			amount:  "123.00",
		},
		{
			name:    "dot and no digits after dot",
			input:   "123.",
			kopecks: 12300,
			amount:  "123.00",
		},
		{
			name:    "dot and no digits before dot",
			input:   ".123",
			kopecks: 12,
			amount:  "0.12",
		},
		{
			name:    "amount zero",
			input:   "0",
			kopecks: 0,
			amount:  "0.0",
		},
		{
			name:    "dot and no digits before and after dot",
			input:   ".",
			kopecks: 0,
			amount:  "0.00",
			err:     ErrInvalidAmount,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kopecks, err := ConvertStringAmountToKopecks(tt.input)
			if tt.err != nil {
				assert.ErrorAs(t, err, &tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.kopecks, kopecks)

				amount := ConvertKopecksToAmount(kopecks)
				assert.Equal(t, tt.amount, amount)

			}
		})
	}
}
