package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	suntime "github.com/alexstojda/fan-light-control/internal"
	"github.com/warthog618/go-gpiocdev"
)

func main() {
	chip := flag.String("chip", "", "GPIO Chip to use. ex: 'gpiochip0'")
	pin := flag.Int("pin", -1, "GPIO pin which the fan is connected to")
	flag.Parse()

	if chip == nil || *chip == "" || pin == nil || *pin < 0 {
		writeErrorMsgAndDie("chip and pin flags are required")
	}

	fanLedPin, err := gpiocdev.RequestLine(*chip, *pin, gpiocdev.AsOutput())
	if err != nil {
		writeErrorMsgAndDie("error opening GPIO pin: %v", err)
	}

	defer fanLedPin.Close()
	
	err = evaluteLightStatus(fanLedPin)
	if err != nil {
		writeErrorMsgAndDie("evaluating light status: %v", err)
	}

}

func evaluteLightStatus(pin *gpiocdev.Line) error {
	times, err := suntime.GetTimes()
	if err != nil {
		return fmt.Errorf("getting sunrise/sunset times: %v", err)
	}

	now := time.Now()

	fmt.Printf("NOW:\t\t%s\nSUNRISE:\t%s\nSUNSET:\t\t%s\n", now.String(), times.Sunrise.String(), times.Sunset.String())

	if now.After(times.Sunrise.Time) && now.Before(times.Sunset.Time) {
		fmt.Println("Turning on")
		pin.SetValue(1)
	} else {
		fmt.Println("Turning off")
		pin.SetValue(0)
	}

	return nil
}

func writeErrorMsgAndDie(msg string, a ...any) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("ERROR: %s\n", msg), a...)
	os.Exit(1)
}
