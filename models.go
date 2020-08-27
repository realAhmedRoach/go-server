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
	UID string
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
	Symbol   string
	Name     string
	FirmType FirmType
}

type Order struct {
	Model
	FirmID    string    `db:"firm_id"`
	OrderTime time.Time `db:"order_time"`
	Price     float64
	Quantity  int
	Side      OrderSide
	OrderType OrderType `db:"order_type"`
}

type SukukOrder struct {
	Order
	Sukuk Sukuk
}

type SalamOrder struct {
	Order
	Commodity Commodity
	Contract  Contract
}

type Transaction struct {
	Model
	LongFirm  Firm      `db:"long_firm"`
	ShortFirm Firm      `db:"short_firm"`
	FilledAt  time.Time `db:"filled_at"`
	Quantity  int
	Price     float64
}

type SukukTransaction struct {
	Transaction
	Sukuk Sukuk
}

type SalamTransaction struct {
	Transaction
	Commodity Commodity
	Contract  Contract
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
	SUKUKORDERSCHEMA = "insert into sukuk_order values ($1, $2, $3, $4, $5, $6) returning uid"
	SALAMORDERSCHEMA = "insert into salam_order values ($1, $2, $3, $4, $5, $6, $7) returning uid"
	SUKUKTRANSSCHEMA = "insert into sukuk_transaction values ($1, $2, $3, $4, $5) returning uid"
	SALAMTRANSSCHEMA = "insert into salam_transaction values ($1, $2, $3, $4, $5, $6) returning uid"
)
