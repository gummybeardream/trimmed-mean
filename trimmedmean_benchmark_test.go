package trimmedmean

import "testing"

func BenchmarkMeanIntegers100(b *testing.B) {
	values := make([]int, 100)

	for index := range values {
		values[index] = (index * 37) % 100
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_, err := Mean(values, 0.05)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMeanFloats100(b *testing.B) {
	values := make([]float64, 100)

	for index := range values {
		values[index] = float64((index*37)%100) + 0.25
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_, err := Mean(values, 0.05)
		if err != nil {
			b.Fatal(err)
		}
	}
}
