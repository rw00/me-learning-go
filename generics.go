package main

import "errors"

type lineItem interface {
	getName() string
	getCost() float64
}

func chargeForLineItem[T lineItem](balance float64, items []T, newItem T) (float64, []T, error) {
	if newItem.getCost() > balance {
		return balance, items, errors.New("insufficient funds")
	}
	return balance - newItem.getCost(), append(items, newItem), nil
}
