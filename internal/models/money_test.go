package models

import (
	"encoding/json"
	"testing"
)

func TestMoneyJSONPreservesDecimalNumber(t *testing.T) {
	var money Money
	if err := json.Unmarshal([]byte(`75000.25`), &money); err != nil {
		t.Fatal(err)
	}

	encoded, err := json.Marshal(money)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != "75000.25" {
		t.Fatalf("expected JSON number 75000.25, got %s", encoded)
	}
}

func TestMoneyScansPostgresNumericBytes(t *testing.T) {
	var money Money
	if err := money.Scan([]byte("75000.00")); err != nil {
		t.Fatal(err)
	}
	if money.String() != "75000" {
		t.Fatalf("expected 75000, got %s", money.String())
	}
}
