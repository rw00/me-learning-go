package main

import (
	"reflect"
	"testing"
)

func TestGetCostsByDay(t *testing.T) {
	costs := []cost{
		{day: 0, value: 10.0},
		{day: 1, value: 20.0},
		{day: 2, value: 30.0},
		{day: 1, value: 15.0},
		{day: 2, value: 25.0},
		{day: 5, value: 13.5},
	}
	expected := []float64{10.0, 35.0, 55.0, 0.0, 0.0, 13.5}

	result := getCostsByDay(costs)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v; got %v", expected, result)
	}
}
