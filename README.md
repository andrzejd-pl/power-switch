# PC power switch

Mini golang package, that simplify using GPIO in Raspberry PI as PC power switch.

## How to use
Example:

```go
package main

import (
	"github.com/andrzejd-pl/power-switch/gpio"
	"log"
	"os"
	"time"
)

const (
	waitTime time.Duration = time.Second / 4
	gpioPin int = 17
	logFlag = 0
	logPrefix = "GPIO"
)

func main() {
	logger := log.New(os.Stdout, logPrefix, logFlag)
	action(logger)
}

func action(logger *log.Logger) {
	powerSwitch, err := gpio.NewGpioPowerSwitch(logger, gpioPin, waitTime)

	if err != nil {
		log.Panicln(err.Error())
	}

	defer powerSwitch.Close()

	log.Println("Turn on PC")
	powerSwitch.TurnOnPc()
}
```

## ToDo
- godocs
- ...