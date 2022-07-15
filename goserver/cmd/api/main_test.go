package main_test

import "testing"

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 5
	if got != expected {
		t.Errorf("error got %v, wanted %v ", got, expected)

	}
}

func TestSubstraction(t *testing.T) {
	got := 3 - 2
	expected := 1
	if got != expected {
		t.Errorf("error got %v, wanted %v ", got, expected)
	}
}
