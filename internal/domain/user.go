package domain

type User struct {
	UserID          int64 `gorm:"primarykey"`
	DefaultCurrency string
}
