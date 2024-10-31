package balance

import (
	"strings"
	"testing"
)

func TestNewBalance(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		wantErr bool
	}{
		{"valid mettler", "mettler", false},
		{"valid sartorius", "sartorius", false},
		{"invalid type", "invalid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := New(tt.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && b == nil {
				t.Errorf("New() returned nil balance for valid format")
			}
		})
	}
}

func TestBalanceGenerateWeight(t *testing.T) {
	b, _ := New("mettler")

	for i := 0; i < 100; i++ {
		weight := b.GenerateWeight()
		if weight < b.Format.MinWeight || weight > b.Format.MaxWeight {
			t.Errorf("GenerateWeight() = %v, want between %v and %v",
				weight, b.Format.MinWeight, b.Format.MaxWeight)
		}
	}
}

func TestBalanceFormatOutput(t *testing.T) {
	b, _ := New("mettler")
	b.current = 123.4567

	output := b.FormatOutput()

	if !strings.HasPrefix(output, b.Format.Prefix) {
		t.Errorf("FormatOutput() prefix = %v, want %v",
			output, b.Format.Prefix)
	}

	if !strings.HasSuffix(output, b.Format.Suffix) {
		t.Errorf("FormatOutput() suffix = %v, want %v",
			output, b.Format.Suffix)
	}
}
