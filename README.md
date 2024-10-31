# Balance Simulator

The Balance Simulator is a tool designed to simulate data from different types of precision balances, such as Mettler and Sartorius balances, via an RS-232 COM port. This simulator can be used for development and testing purposes where a physical balance is unavailable.

## Features

- Simulates serial data output from precision balances
- Supports multiple balance types: Mettler and Sartorius
- Configurable data transmission interval

## Prerequisites

To use the Balance Simulator, ensure you have the following:

## Installation

- Go installed (version 1.18+)
- Required serial port setup (e.g., `COM1`)


Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/thisishugow/balance-sim.git
cd balance-sim
go get
```

## Usage

The Balance Simulator supports various configurations via command-line flags. The available flags include the COM port name, balance type, and data transmission interval.
- `-port`: Specifies the serial port name (default: COM1)
- `-type`: Specifies the type of balance to simulate. Supported values are mettler and sartorius (default: mettler)
- `-interval`: Sets the interval for sending data in milliseconds (default: 1000)
```bash
go run ./cmd/simulator/main.go -port COM1 -type mettler -interval 1000
```


## Structure

The recommended structure for this project is as follows:

* `cmd/` - contains the main application code
	+ `simulator/` - contains the simulator code
	+ `main.go` - the main entry point
* `internal/` - contains the internal code
	+ `balance/` - contains the balance code
		- `balance.go` - the balance struct and methods
		- `balance_test.go` - tests for the balance code
	+ `serial/` - contains the serial port code
		- `port.go` - the port struct and methods
		- `port_test.go` - tests for the serial port code
	+ `simulator/` - contains the simulator code
		- `simulator.go` - the simulator struct and methods
		- `simulator_test.go` - tests for the simulator code
* `go.mod` - the Go module file
* `go.sum` - the Go module checksums file

