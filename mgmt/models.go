package mgmt

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
	OrderSide OrderSide `json:"side,omitempty"`
	OrderType OrderType `json:"type,omitempty"`
}

type SukukOrder struct {
	Order
	Sukuk Sukuk `json:"Sukuk,omitempty"`
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

type Service interface {
	Get(uid string) (string, error)
	Put(values ...interface{}) (string, error)
	Delete(uid string) error
}

type SukukTransaction struct {
	Transaction
	Sukuk Sukuk `json:"Sukuk,omitempty"`
}

type SalamTransaction struct {
	Transaction
	Commodity Commodity `json:"commodity,omitempty"`
	Contract  Contract  `json:"contract,omitempty"`
}

const TESTFIRMID = "baf78936-5986-4f24-8a40-5e11aef970c6"
