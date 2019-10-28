package main

import "testing"

func TestThreeSum(t *testing.T) {
	// Should count the number of triples in a file of N integers that sum to 0
	arr := []int{5, 9, -6, 2, 4}
	count := ThreeSum(arr)
	if count != 1 {
		t.Errorf("ThreeSum was incorrect, got: %d, want: %d", count, 1)
	}
}

func TestThreeSumNoResults(t *testing.T) {
	// Should count the number of triples in a file of N integers that sum to 0
	arr := []int{5, 9, -3, 2, 4}
	count := ThreeSum(arr)
	if count != 0 {
		t.Errorf("Passed a slice with no possible results be we got: %d, want: %d", count, 0)
	}
}
