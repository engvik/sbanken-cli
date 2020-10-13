package sbanken

import (
	"context"

	"github.com/engvik/sbanken-go"
)

type testClient struct{}

func (c testClient) ListCards(context.Context) ([]sbanken.Card, error) { return nil, nil }
func (c testClient) ListEfakturas(context.Context, *sbanken.EfakturaListQuery) ([]sbanken.Efaktura, error) {
	return nil, nil
}
func (c testClient) PayEfaktura(context.Context, *sbanken.EfakturaPayQuery) error { return nil }
func (c testClient) ListNewEfakturas(context.Context, *sbanken.EfakturaListQuery) ([]sbanken.Efaktura, error) {
	return nil, nil
}
func (c testClient) ReadEfaktura(context.Context, string) (sbanken.Efaktura, error) {
	return sbanken.Efaktura{}, nil
}
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
