// Package trimmedmean calculates means after removing observations from the lower and upper ends of a sorted dataset
package trimmedmean

import (
	"errors"
	"math"
	"slices"
)

// Number represents the numeric types accepted by this package.
type Number interface {
	~int | ~float64
}

// Mean calculates the trimmed mean of values
//
// One trim proportion applies symmetric trimming. Two proportions apply
// separate lower and upper trimming.
func Mean[T Number](values []T, trims ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("values must not be empty")
	}

	if len(trims) < 1 || len(trims) > 2 {
		return 0, errors.New("provide one or two trim proportions")
	}

	lowerTrim := trims[0]
	upperTrim := lowerTrim

	if len(trims) == 2 {
		upperTrim = trims[1]
	}

	if math.IsNaN(lowerTrim) || math.IsNaN(upperTrim) {
		return 0, errors.New("trim proportions must be numbers")
	}

	if lowerTrim < 0 || lowerTrim >= 1 || upperTrim < 0 || upperTrim >= 1 {
		return 0, errors.New("trim proportions must be at least 0 and less than 1")
	}

	if lowerTrim+upperTrim >= 1 {
		return 0, errors.New("combined trim proportions must be less than 1")
	}

	sortedValues := slices.Clone(values)
	slices.Sort(sortedValues)

	lowerCount := int(float64(len(sortedValues)) * lowerTrim)
	upperCount := int(float64(len(sortedValues)) * upperTrim)

	keptValues := sortedValues[lowerCount : len(sortedValues)-upperCount]

	var sum float64
	for _, value := range keptValues {
		sum += float64(value)
	}

	return sum / float64(len(keptValues)), nil
}
