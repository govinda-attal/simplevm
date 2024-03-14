package accounts

import (
	"time"

	"github.com/govinda-attal/simple-vm/types/ordmap"
)

type Property struct {
	Type string
	Name string
}

type PropertyValue string

type CurrentProperties map[Property]PropertyValue

type OrderedPropertyValues ordmap.OrderedMap[time.Time, PropertyValue]

type Properties map[Property]OrderedPropertyValues
