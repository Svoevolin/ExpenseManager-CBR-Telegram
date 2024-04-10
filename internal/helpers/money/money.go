package money

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

var ErrInvalidAmount = errors.New("invalid amount")

// 1,000,500.10 -> 1000500.10 || 1 000 500.100 -> 1000500.100
var regexpNoDigit = regexp.MustCompile(`[^\d\.]`)

func ConvertStringAmountToKopecks(amount string) (int64, error) {
	v, err := strconv.ParseFloat(regexpNoDigit.ReplaceAllString(amount, ""), 64)
	if err != nil {
		return 0, ErrInvalidAmount
	}
	return ConvertFloat64AmountToKopecks(v)
}

func ConvertFloat64AmountToKopecks(amount float64) (int64, error) {
	return int64(100 * amount), nil
}

func ConvertKopecksToAmount(kopecks int64) string {
	amount := fmt.Sprintf("%d", kopecks)
	if len(amount) < 3 {
		return fmt.Sprintf("0.%s", amount)
	}
	return fmt.Sprintf("%s.%s", amount[:len(amount)-2], amount[len(amount)-2:])
}
