package mypackage

import "testing"

func TestAverage(t *testing.T) {
	v, err := Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got", v)
	}

	v, err = Average([]float64{})
	if v != 0 || err != nil {
		t.Error("For `v, err := Average([]float64{})`",
			"Expected: 0, nil",
			"but got: ", v, err)
	}
}

func TestMin(t *testing.T) {
	v := Min([]float64{10, 20, 30, 90, 1, 2, 3, 44, 0, 1, 12, 0, 21})
	if v != 0 {
		t.Error("Expected 0, got", v)
	}
}

func TestMax(t *testing.T) {
	v := Max([]float64{10, 20, 30, 90, 1, 2, 3, 44, 0, 1, 12, 0, 21})
	if v != 90 {
		t.Error("Expected 90, got", v)
	}
}

