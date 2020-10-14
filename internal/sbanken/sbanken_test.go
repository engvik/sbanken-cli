package sbanken

import (
	"context"

	"github.com/engvik/sbanken-go"
)

type testClient struct{}

func (c testClient) Transfer(context.Context, *sbanken.TransferQuery) error { return nil }
