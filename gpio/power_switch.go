package gpio

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

const (
	WaitTime time.Duration = time.Second / 4
)

func TurnOnPc(pin rpio.Pin) {
	pin.High()
	time.Sleep(WaitTime)
	pin.Low()
}
