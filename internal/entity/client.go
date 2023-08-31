package entity

import (
	"errors"
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
	client := &Client{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("invalid name")
	}
	if c.Email == "" {
		return errors.New("invalid email")
	}
	return nil
}
