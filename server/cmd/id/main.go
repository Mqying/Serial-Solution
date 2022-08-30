package main

import (
	"fmt"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/serial"
)

var machineType = model.AcidType

func main() {
	v := &serial.VoltageSensor{}

	var err error

	if err = v.OpenSerial(machineType); err != nil {
		fmt.Println("Open Serial Error:", err)
		return
	}

	if s, err := v.GetDeviceID(machineType); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%v\n", s)
	}
}
