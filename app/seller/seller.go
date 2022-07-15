package seller

import "time"

type Seller struct {
	ID       int
	Name     string
	Gender   string
	ItemID   int
	CreateAt time.Time
	UpdateAt time.Time
}
