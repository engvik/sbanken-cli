package sbanken

import (
	"context"

	"github.com/engvik/sbanken-go"
)

type testClient struct{}

func (c testClient) ListPayments(context.Context, string, *sbanken.PaymentListQuery) ([]sbanken.Payment, error) {
	return nil, nil
}
func (c testClient) ReadPayment(context.Context, string, string) (sbanken.Payment, error) {
	return sbanken.Payment{}, nil
}
func (c testClient) ListStandingOrders(context.Context, string) ([]sbanken.StandingOrder, error) {
	return nil, nil
}
func (c testClient) ListTransactions(context.Context, string, *sbanken.TransactionListQuery) ([]sbanken.Transaction, error) {
	return nil, nil
}
func (c testClient) Transfer(context.Context, *sbanken.TransferQuery) error { return nil }
