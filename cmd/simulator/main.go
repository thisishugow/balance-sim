package main

import (
	"flag"
	"log"
	"time"

	"balance-sim/internal/simulator"
)

// main starts a balance simulator with the given configuration.
//
// The configuration is specified through the following command line flags:
//
// -port: the name of the serial port to use (default: COM1)
// -type: the type of balance to simulate (default: mettler, supported values: mettler, sartorius)
// -interval: the interval at which data is sent in milliseconds (default: 1000)
//
// The simulator will continuously send data to the given serial port at the given interval.
// The program will exit when the simulator is stopped.
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
