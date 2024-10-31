package serial

import (
	"testing"
)

func TestNewPort(t *testing.T) {
	// 注意：這個測試在沒有實際 COM 端口時會失敗
	t.Skip("Skipping serial port test - requires hardware")

	_, err := New("COM1")
	if err != nil {
		t.Errorf("New() error = %v", err)
	}
}
