package simulator

import (
	"balance-sim/internal/balance"
	"balance-sim/internal/serial"
	"log"
	"time"
)

type Config struct {
	PortName    string
	Interval    time.Duration
	BalanceType string
}

type Simulator struct {
	balance *balance.Balance
	port    *serial.Port
	stop    chan struct{}
	config  Config
}

// New returns a new Simulator based on the given config.
// It returns an error if there's a problem initializing the balance or port.
func New(config Config) (*Simulator, error) {
	bal, err := balance.New(config.BalanceType)
	if err != nil {
		return nil, err
	}

	port, err := serial.New(config.PortName)
	if err != nil {
		return nil, err
	}

	return &Simulator{
		balance: bal,
		port:    port,
		stop:    make(chan struct{}),
		config:  config,
	}, nil
}

// Start starts the simulator, sending data to the port at the given interval.
// It runs until Stop is called.
func (s *Simulator) Start() {
	ticker := time.NewTicker(s.config.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.balance.GenerateWeight()
			data := s.balance.FormatOutput()

			if err := s.port.Write(data); err != nil {
				log.Printf("發送數據錯誤: %v", err)
				continue
			}

			log.Printf("已發送: %s", data)

		case <-s.stop:
			return
		}
	}
}

// Stop stops the simulator, causing it to stop sending data to the port.
func (s *Simulator) Stop() {
	close(s.stop)
}

// Close stops the simulator and closes the serial port.
// It's a good idea to call this when you're done using the simulator, to free up system resources.
func (s *Simulator) Close() error {
	s.Stop()
	return s.port.Close()
}
