package models

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type Money struct {
	decimal.Decimal
}

func (m *Money) UnmarshalJSON(data []byte) error {
	return m.Decimal.UnmarshalJSON(data)
}

func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(m.String()), nil
}

var _ json.Marshaler = Money{}
var _ json.Unmarshaler = (*Money)(nil)
