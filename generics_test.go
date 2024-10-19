package main

import (
	"errors"
	"testing"
)

type mockLineItem struct {
	name string
	cost float64
}

func TestChargeForLineItem(t *testing.T) {
	tests := []struct {
		name        string
		balance     float64
		items       []mockLineItem
		newItem     mockLineItem
		wantBalance float64
		wantErr     error
	}{
		{
			name:        "sufficient funds",
			balance:     100.0,
			items:       []mockLineItem{},
			newItem:     mockLineItem{name: "item1", cost: 50.0},
			wantBalance: 50.0,
			wantErr:     nil,
		},
		{
			name:        "insufficient funds",
			balance:     30.0,
			items:       []mockLineItem{},
			newItem:     mockLineItem{name: "item2", cost: 50.0},
			wantBalance: 30.0,
			wantErr:     errors.New("insufficient funds"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBalance, gotItems, gotErr := chargeForLineItem(tt.balance, tt.items, tt.newItem)
			if gotBalance != tt.wantBalance {
				t.Errorf("chargeForLineItem() got balance = %v, want = %v", gotBalance, tt.wantBalance)
			}
			if gotErr != nil && gotErr.Error() != tt.wantErr.Error() {
				t.Errorf("chargeForLineItem() got err = %v, want = %v", gotErr, tt.wantErr)
			}
			if len(gotItems) != len(tt.items)+1 && tt.wantErr == nil {
				t.Errorf("chargeForLineItem() got items = %v, want = %v", gotItems, tt.items)
			}
		})
	}
}

func (m mockLineItem) getName() string {
	return m.name
}

func (m mockLineItem) getCost() float64 {
	return m.cost
}
