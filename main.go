package main

import (
	"fmt"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("failed to rpio.Open()")
		fmt.Println(err.Error())
	}
	inputpin := rpio.Pin(27)
	// outputpin := rpio.Pin(17)

	inputpin.Input()
	// outputpin.Output()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(inputpin.Read())
		// if inputpin.Read() == rpio.High {
		// fmt.Println("High")
		// outputpin.High()
		// } else if inputpin.Read() == rpio.Low {
		// outputpin.Low()
		// }
	}
}
