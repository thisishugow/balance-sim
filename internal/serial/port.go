package serial

import (
	"time"

	"github.com/tarm/serial"
)

type Port struct {
	port *serial.Port
}

// New opens a serial port with the given name and a default configuration.
//
// It returns the new Port and an error if the port can't be opened.
func New(portName string) (*Port, error) {
	config := &serial.Config{
		Name:        portName,
		Baud:        9600,
		Size:        serial.DefaultSize,
		Parity:      serial.ParityNone,
		StopBits:    serial.Stop1,
		ReadTimeout: time.Second * 1,
	}

	port, err := serial.OpenPort(config)
	if err != nil {
		return nil, err
	}

	return &Port{port: port}, nil
}

// Write writes the given string to the serial port.
//
// It returns an error if there's a problem writing to the port.
func (p *Port) Write(data string) error {
	_, err := p.port.Write([]byte(data))
	return err
}

// Close closes the serial port.
//
// It's a good idea to call this when you're done using the port, to free up system
// resources.
func (p *Port) Close() error {
	return p.port.Close()
}
