package entity

import (
	"time"
)

type Photo struct {
	ImageUrl     string
	ShootingDate time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewPhoto(image string, shootingTime, createdAt, updatedAt time.Time) *Photo {
	return &Photo{image, shootingTime, createdAt, updatedAt}
}
