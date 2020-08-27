package main

import (
	"time"
)

type Sukuk string
type Commodity string
type Contract string

type OrderSide string

const (
	BUY  OrderSide = "BUY"
	SELL           = "SELL"
)

type OrderType string

const (
	MKT OrderType = "MARKET"
	LMT           = "LIMIT"
)

type FirmType string

const (
	TRDR FirmType = "TRDR"
	WRHS          = "WRHS"
	PROD          = "PROD"
	CONS          = "CONS"
	MKMK          = "MKMK"
)

type Model struct {
	UID string `json:"uid,omitempty"`
}

type DatabaseModel interface {
	TableName() string
}

const (
	DB_FIRM             = "firm"
	DB_SUKUKORDER       = "sukuk_order"
	DB_SALAMORDER       = "salam_order"
	DB_SUKUKTRANSACTION = "sukuk_transaction"
	DB_SALAMTRANSACTION = "salam_transaction"
)

type Firm struct {
	Model
	Symbol   string   `json:"symbol,omitempty"`
	Name     string   `json:"name,omitempty"`
	FirmType FirmType `json:"firm_type,omitempty"`
}

type Order struct {
	Model
	FirmID    string    `json:"firm_id,omitempty"`
	OrderTime time.Time `json:"order_time,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Quantity  int       `json:"quantity,omitempty"`
	Side      OrderSide `json:"side,omitempty"`
	OrderType OrderType `json:"order_type,omitempty"`
}

type SukukOrder struct {
	Order
	Sukuk Sukuk `json:"sukuk,omitempty"`
}

type SalamOrder struct {
	Order
	Commodity Commodity `json:"commodity,omitempty"`
	Contract  Contract  `json:"contract,omitempty"`
}

type Transaction struct {
	Model
	LongFirm  Firm      `json:"long_firm,omitempty"`
	ShortFirm Firm      `json:"short_firm,omitempty"`
	FilledAt  time.Time `json:"filled_at,omitempty"`
	Quantity  int       `json:"quantity,omitempty"`
	Price     float64   `json:"price,omitempty"`
}

type SukukTransaction struct {
	Transaction
	Sukuk Sukuk `json:"sukuk,omitempty"`
}

type SalamTransaction struct {
	Transaction
	Commodity Commodity `json:"commodity,omitempty"`
	Contract  Contract  `json:"contract,omitempty"`
}

func (firm *Firm) TableName() string {
	return DB_FIRM
}

func (*SukukOrder) TableName() string {
	return DB_SUKUKORDER
}

func (*SalamOrder) TableName() string {
	return DB_SALAMORDER
}

func (*SalamTransaction) TableName() string {
	return DB_SALAMTRANSACTION
}

func (*SukukTransaction) TableName() string {
	return DB_SUKUKTRANSACTION
}

const (
	SUKUKORDERSCHEMA = `
		insert into sukuk_order 
			(firm_id, sukuk, price, quantity, side, order_type) 
			values ($1, $2, $3, $4, $5, $6) 
			returning uid`
	SALAMORDERSCHEMA = "insert into salam_order values ($1, $2, $3, $4, $5, $6, $7) returning uid"
	SUKUKTRANSSCHEMA = "insert into sukuk_transaction values ($1, $2, $3, $4, $5) returning uid"
	SALAMTRANSSCHEMA = "insert into salam_transaction values ($1, $2, $3, $4, $5, $6) returning uid"
)
