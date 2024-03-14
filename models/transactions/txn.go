package transactions

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/govinda-attal/simple-vm/types/ordmap"
)

type balancedTx struct {
	Id   string
	ty   string
	from []Unit
	to   []Unit
}

type Unit struct {
	Type     string
	Currency string
	Amount   *money.Money
}

type StdOutgoing balancedTx

type StdIncoming balancedTx

type Transaction balancedTx

type OrderedTransactions ordmap.OrderedMap[time.Time, Transaction]
