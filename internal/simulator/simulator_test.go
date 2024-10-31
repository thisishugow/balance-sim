package simulator

import (
	"testing"
	"time"
)

func TestNewSimulator(t *testing.T) {
	// 注意：這個測試在沒有實際 COM 端口時會失敗
	t.Skip("Skipping simulator test - requires hardware")

	config := Config{
		PortName:    "COM1",
		Interval:    time.Second,
		BalanceType: "mettler",
	}

	sim, err := New(config)
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}
	defer sim.Close()

	if sim.balance == nil {
		t.Error("New() balance is nil")
	}
	if sim.port == nil {
		t.Error("New() port is nil")
	}
}
