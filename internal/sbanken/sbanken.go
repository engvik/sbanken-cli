package sbanken

import (
	"context"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
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

func (c *Connection) ConnectClient(ctx context.Context, cliCtx *cli.Context) error {
	cfg := &sbanken.Config{
		ClientID:     cliCtx.String("client-id"),
		ClientSecret: cliCtx.String("client-secret"),
		CustomerID:   cliCtx.String("customer-id"),
	}
	sClient, err := sbanken.NewClient(ctx, cfg, nil)
	if err != nil {
		return err
	}

	c.Client = sClient

	return nil
}
