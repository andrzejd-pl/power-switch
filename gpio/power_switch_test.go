package gpio

import (
	"github.com/stianeikeland/go-rpio"
	"testing"
)

const (
	SrcGpioPin  int = 4
	TestGpioPin int = 3
)

func TestTurnOnPc(t *testing.T) {
	err := rpio.Open()

	if err != nil {
		t.Errorf("Error with initialize GPIO: %s", err.Error())
	}

	defer rpio.Close()

	src := rpio.Pin(SrcGpioPin)
	test := rpio.Pin(TestGpioPin)
	src.Mode(rpio.Output)
	test.Mode(rpio.Input)
	test.Detect(rpio.RiseEdge)

	TurnOnPc(src)

	if !test.EdgeDetected() {
		t.Errorf("Error with turning pc")
	}
}
