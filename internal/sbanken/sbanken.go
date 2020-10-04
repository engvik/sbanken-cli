package sbanken

import (
	"context"

	"github.com/engvik/sbanken-go"
)

type Connection struct {
	Client *sbanken.Client
}

type Config struct {
	ClientID     string
	ClientSecret string
	CustomerID   string
}

func NewEmptyConnection() *Connection {
	return &Connection{}
}

func (c *Connection) ConnectClient(ctx context.Context, cfg *Config) error {
	config := &sbanken.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		CustomerID:   cfg.CustomerID,
	}

	sClient, err := sbanken.NewClient(ctx, config, nil)
	if err != nil {
		return err
	}

	c.Client = sClient

	return nil
}
