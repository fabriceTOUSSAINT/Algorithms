package main

import "testing"

func TestTwoSum(t *testing.T) {
	// Should count the number of triples in a file of N integers that sum to 0
	arr := []int{5, 9, -6, 6, 4}
	count := TwoSum(arr)
	if count != 1 {
		t.Errorf("TwoSum was incorrect, got: %d, want: %d", count, 1)
	}
}

func TestTwoSumNoResults(t *testing.T) {
	// Should count the number of triples in a file of N integers that sum to 0
	arr := []int{5, 9, -3, 2, 4}
	count := TwoSum(arr)
	if count != 0 {
		t.Errorf("Passed a slice with no possible results be we got: %d, want: %d", count, 0)
	}
}
