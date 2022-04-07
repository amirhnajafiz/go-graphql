package store

import "time"

type Book struct {
	Title       string
	Reference   string
	PublishDate time.Time
	Sells       int
	Price       float64
}
