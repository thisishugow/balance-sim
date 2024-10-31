# Balance Simulator

This is a simple simulator for electronic balances. It can be used to generate random weights at a specified interval.

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

## Usage

To use the simulator, run the following command:
```bash
go run cmd/simulator/main.go
```