package gpio

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
)

type PowerSwitch interface {
	TurnOnPc()
	Close() error
}

type gpioPowerSwitch struct {
	logger   *log.Logger
	pin      rpio.Pin
	waitTime time.Duration
}

func NewGpioPowerSwitch(logger *log.Logger, pinNumber int, waitTime time.Duration) (PowerSwitch, error) {
	err := rpio.Open()
	if err != nil {
		return nil, err
	}

	pin := rpio.Pin(pinNumber)
	pin.Output()

	return gpioPowerSwitch{
		logger:   logger,
		pin:      pin,
		waitTime: waitTime,
	}, nil
}

func (powerSwitch gpioPowerSwitch) TurnOnPc() {
	powerSwitch.pin.High()
	time.Sleep(powerSwitch.waitTime)
	powerSwitch.pin.Low()
}

func (powerSwitch gpioPowerSwitch) Close() error {
	err := rpio.Close()

	if err != nil {
		powerSwitch.logger.Println("error: ", err.Error())
	}

	return err
}
