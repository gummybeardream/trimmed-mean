package trimmedmean

import (
	"math"
	"reflect"
	"testing"
)

func TestMeanSymmetricIntegers(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got, err := Mean(values, 0.10)
	if err != nil {
		t.Fatalf("Mean returned an unexpected error: %v", err)
	}

	want := 5.5
	if !floatsAlmostEqual(got, want) {
		t.Errorf("Mean() = %v; want %v", got, want)
	}
}

func TestMeanEmptySlice(t *testing.T) {
	values := []float64{}

	_, err := Mean(values, 0.05)
	if err == nil {
		t.Error("Mean returned no error for an empty slice")
	}
}

func TestMeanSymmetricFloats(t *testing.T) {
	values := []float64{1.5, 2.5, 3.5, 4.5, 100.5}

	got, err := Mean(values, 0.20)
	if err != nil {
		t.Fatalf("Mean returned an unexpected error: %v", err)
	}

	want := 3.5
	if !floatsAlmostEqual(got, want) {
		t.Errorf("Mean() = %v; want %v", got, want)
	}
}

func TestMeanAsymmetric(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got, err := Mean(values, 0.10, 0.20)
	if err != nil {
		t.Fatalf("Mean returned an unexpected error: %v", err)
	}

	want := 5.0
	if !floatsAlmostEqual(got, want) {
		t.Errorf("Mean() = %v; want %v", got, want)
	}
}

func TestMeanZeroTrimUnsorted(t *testing.T) {
	values := []int{5, 1, 3, 2, 4}

	got, err := Mean(values, 0)
	if err != nil {
		t.Fatalf("Mean returned an unexpected error: %v", err)
	}

	want := 3.0
	if !floatsAlmostEqual(got, want) {
		t.Errorf("Mean() = %v; want %v", got, want)
	}
}

func TestMeanDoesNotModifyInput(t *testing.T) {
	values := []int{5, 1, 4, 2, 3}
	original := slicesClone(values)

	_, err := Mean(values, 0.20)
	if err != nil {
		t.Fatalf("Mean returned an unexpected error: %v", err)
	}

	if !reflect.DeepEqual(values, original) {
		t.Errorf("Mean modified the input: got %v; want %v", values, original)
	}
}

func slicesClone[T any](values []T) []T {
	cloned := make([]T, len(values))
	copy(cloned, values)
	return cloned
}

func TestMeanInvalidTrims(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}

	tests := []struct {
		name  string
		trims []float64
	}{
		{
			name:  "no trim arguments",
			trims: nil,
		},
		{
			name:  "more than two trim arguments",
			trims: []float64{0.10, 0.10, 0.10},
		},
		{
			name:  "negative lower trim",
			trims: []float64{-0.10},
		},
		{
			name:  "negative upper trim",
			trims: []float64{0.10, -0.10},
		},
		{
			name:  "lower trim equal to one",
			trims: []float64{1.0, 0},
		},
		{
			name:  "upper trim equal to one",
			trims: []float64{0, 1.0},
		},
		{
			name:  "combined trims equal to one",
			trims: []float64{0.50, 0.50},
		},
		{
			name:  "combined trims greater than one",
			trims: []float64{0.60, 0.50},
		},
		{
			name:  "lower trim is NaN",
			trims: []float64{math.NaN()},
		},
		{
			name:  "upper trim is NaN",
			trims: []float64{0.10, math.NaN()},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Mean(values, test.trims...)
			if err == nil {
				t.Errorf("Mean with trims %v returned no error", test.trims)
			}
		})
	}
}

func TestMeanFractionalTrimCountRoundsDown(t *testing.T) {
	values := []int{
		0, 1, 2, 3, 4,
		5, 6, 7, 8, 9,
		10, 11, 12, 100, 1000,
	}

	got, err := Mean(values, 0.10)
	if err != nil {
		t.Fatalf("Mean returned an unexpected error: %v", err)
	}

	want := 178.0 / 13.0
	if !floatsAlmostEqual(got, want) {
		t.Errorf("Mean() = %v; want %v", got, want)
	}
}

func floatsAlmostEqual(first, second float64) bool {
	const tolerance = 1e-12
	return math.Abs(first-second) <= tolerance
}

func TestMeanRejectsNonFiniteValues(t *testing.T) {
	tests := []struct {
		name   string
		values []float64
	}{
		{
			name:   "NaN observation",
			values: []float64{1, 2, math.NaN(), 4, 5},
		},
		{
			name:   "positive infinity",
			values: []float64{1, 2, math.Inf(1), 4, 5},
		},
		{
			name:   "negative infinity",
			values: []float64{1, 2, math.Inf(-1), 4, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Mean(test.values, 0)
			if err == nil {
				t.Errorf("Mean returned no error for values %v", test.values)
			}
		})
	}
}
