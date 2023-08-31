package entity

import (
	"github.com/google/uuid"
	"time"
)

type Client struct {
	ID       string
	Name     string
	Email    string
	CreateAt time.Time
	UpdateAt time.Time
}

func NewClient(name, email string) (*Client, error) {
	return &Client{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}, nil
}
