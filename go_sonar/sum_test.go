package main

import "testing"

func TesteSum(t *testing.T) {
	result := sum(2, 2)
	if result != 5 {
		t.Error("The result must be 5")
	}
}
