package accounts

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/govinda-attal/simple-vm/types/ordmap"
)

type BalanceType struct {
	Type     string
	Currency string
}

type Balance struct {
	Amount    *money.Money
	TxEntryID string
}

type OrderedBalance ordmap.OrderedMap[time.Time, Balance]

type CurrentBalances map[BalanceType]Balance

type TypeBalances map[BalanceType]OrderedBalance
