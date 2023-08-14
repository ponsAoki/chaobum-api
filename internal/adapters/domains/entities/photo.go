package entity

import (
	"time"
)

type Photo struct {
	ID           string    `json:"id"`
	ImageUrl     string    `json:"imageUrl"`
	ShootingDate string    `json:"shootingDate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func NewPhoto(id string, image, shootingTime string, createdAt, updatedAt time.Time) *Photo {
	return &Photo{id, image, shootingTime, createdAt, updatedAt}
}
