package buku

import "time"

type Buku struct {
	ID          int
	Title       string
	Description string
	Price       int16
	Rating      int16
	CreatedAt   time.Time
	UpdatedAt   time.Time
	
}
