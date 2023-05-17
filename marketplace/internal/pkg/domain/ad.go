package domain

import (
	"time"
)

type Ad struct {
	id          string
	Title       string
	Description string
	Price       float32
	date        time.Time
}

func NewAd(title, description string, price float32) Ad {
	return Ad{
		Title:       title,
		Description: description,
		Price:       price,
		date:        time.Now(),
	}
}

func (ad *Ad) GetId() string {
	return ad.id
}

func (ad *Ad) SetID(id string) {
	ad.id = id
}
