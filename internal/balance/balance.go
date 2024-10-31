package balance

import (
	"fmt"
	"math/rand"
)

type Balance struct {
	Format  Format
	current float64
}

// New returns a new Balance based on the given formatType.
// If the type is not supported, it returns an error.
func New(formatType string) (*Balance, error) {
	format, exists := Formats[formatType]
	if !exists {
		return nil, fmt.Errorf("不支援的天秤類型: %s", formatType)
	}

	return &Balance{
		Format: format,
	}, nil
}

// GenerateWeight generates a random weight between b.Format.MinWeight and
// b.Format.MaxWeight, rounding it to b.Format.Precision decimal places.
// It stores the generated weight in b.current and returns it.
func (b *Balance) GenerateWeight() float64 {
	weight := b.Format.MinWeight + rand.Float64()*(b.Format.MaxWeight-b.Format.MinWeight)
	factor := float64(1)
	for i := 0; i < b.Format.Precision; i++ {
		factor *= 10
	}
	b.current = float64(int(weight*factor)) / factor
	return b.current
}

// FormatOutput returns a formatted string representation of the current weight
// using the balance's format. The result includes the specified prefix and suffix,
// with the weight rounded to the defined precision.
func (b *Balance) FormatOutput() string {
	return fmt.Sprintf("%s%.*f%s",
		b.Format.Prefix,
		b.Format.Precision,
		b.current,
		b.Format.Suffix)
}
