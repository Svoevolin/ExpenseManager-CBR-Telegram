package domain

import "time"

type Expense struct {
	Title  string
	Date   time.Time
	Amount int64
}
