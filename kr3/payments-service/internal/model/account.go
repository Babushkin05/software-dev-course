package model

import "time"

type Account struct {
	ID        string // UUID
	UserID    string // Внешний ID пользователя
	Balance   int64  // В копейках
	CreatedAt time.Time
}
