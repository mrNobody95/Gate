package api

import "github.com/mrNobody95/Gate/models"

type OrderBookResponse struct {
	Symbol string
	Bids   []models.Order
	Asks   []models.Order
}

type RecentTradesResponse struct {
	Symbol string
	Trades []models.Trade
}

type MarketStatusResponse struct {
	Symbol string
	Trades []models.Trade
}

type OHLCResponse struct {
	Symbol     string
	Resolution *models.Resolution
	Candles    []models.Candle
}

type UserInfoResponse struct {
	User        models.User
	BankAccount []models.BankAccount
}

type WalletResponse struct {
	Wallet models.Wallet
}

type WalletsResponse struct {
	Wallets []models.Wallet
}

type BalanceResponse struct {
	Symbol  string
	Balance string
}

type TransactionListResponse struct {
	Transactions []models.Transaction
}

type OrderResponse struct {
	Order models.Order
}

type OrderListResponse struct {
	Orders []models.Order
}

type UpdateOrderStatusResponse struct {
	NewStatus models.OrderStatus
}
