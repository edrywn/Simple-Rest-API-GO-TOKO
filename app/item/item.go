package item

import "time"

type Item struct {
	ID       int
	Name     string
	Price    int
	Stock    int
	CreateAt time.Time
	UpdateAt time.Time
}

