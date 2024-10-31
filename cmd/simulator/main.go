package main

import (
	"flag"
	"log"
	"time"

	"balance-sim/internal/simulator"
)

func main() {
	portName := flag.String("port", "COM1", "COM port name")
	balanceType := flag.String("type", "mettler", "Balance type (mettler/sartorius)")
	interval := flag.Int("interval", 1000, "Data sending interval in milliseconds")
	flag.Parse()

	config := simulator.Config{
		PortName:    *portName,
		Interval:    time.Duration(*interval) * time.Millisecond,
		BalanceType: *balanceType,
	}

	sim, err := simulator.New(config)
	if err != nil {
		log.Fatalf("初始化模擬器失敗: %v", err)
	}
	defer sim.Close()

	log.Printf("模擬器已啟動 - Port: %s, 天秤類型: %s", *portName, *balanceType)
	sim.Start()
}
