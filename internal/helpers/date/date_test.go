package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDate(t *testing.T) {
	t.Run("method should give only the day month and year", func(t *testing.T) {
		now := time.Now()
		date := GetDate(now)
		expect := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		assert.EqualValues(t, date, expect)
	})
}
