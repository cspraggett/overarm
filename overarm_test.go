package main

import (
	"reflect"
	"testing"
)

func TestCalculateRings(t *testing.T) {
	tests := []struct {
		remaingSpace       float64
		ringSize           float64
		wantRemainingSpace float64
		wantNumberOfRings  int
	}{
		{remaingSpace: (8.25 + sixteenth), ringSize: 1.0,
			wantRemainingSpace: float64(0.25 + sixteenth), wantNumberOfRings: 8},
		{remaingSpace: (0.25 + sixteenth), ringSize: quarter,
			wantRemainingSpace: float64(sixteenth), wantNumberOfRings: 1},
		{remaingSpace: (thirtySecond), ringSize: quarter,
			wantRemainingSpace: thirtySecond, wantNumberOfRings: 0},
	}

	for _, tc := range tests {
		spaceLeft, numberOfRings := calculateRings(tc.remaingSpace, tc.ringSize)
		if spaceLeft != tc.wantRemainingSpace || numberOfRings != tc.wantNumberOfRings {
			t.Errorf("got spaceLeft: %f, wanted spaceLeft: %f\ngot numberOfRings: %d, wanted numberOfRings: %d",
				spaceLeft, tc.wantRemainingSpace, numberOfRings, tc.wantNumberOfRings)
		}
	}
}

func TestGetResults(t *testing.T) {
	tests := []struct {
		cutSize float64
		totals  []int
	}{
		{cutSize: (8.0 + sixteenth), totals: []int{8, 0, 0, 0, 1, 0}},
		{cutSize: (8.25 + sixteenth), totals: []int{8, 0, 1, 0, 1, 0}},
		{cutSize: (8.385 + sixteenth), totals: []int{8, 0, 1, 1, 1, 1}},
	}

	for _, tc := range tests {
		gotTotals := getResults(sizes, tc.cutSize)
		if !reflect.DeepEqual(gotTotals, tc.totals) {
			t.Errorf("got %v, want %v", gotTotals, tc.totals)
		}
	}

}

// 8.4475
