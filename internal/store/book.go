package store

import "time"

type Book struct {
	Id          int
	Title       string
	Reference   string
	PublishDate time.Time
	Sells       int
	Price       float64
}
