package models

import "time"

type Chef struct {
	Id         string    `json:"id"`
	Name       string    `json:"name" binding:"required"`
	Country    string    `json:"country"`
	Experience int       `json:"experience"`
	CreatedAt  time.Time `json:"created_at"`
}
