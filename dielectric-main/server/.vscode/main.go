package main

import (
	"fmt"
	"time"

	serial "github.com/silverswords/dielectric/server/serial"
	"github.com/silverswords/dielectric/server/zlog"
	sensor "github.com/tarm/serial"
)

func main() {
	ports := serial.GetPorts()
	fmt.Println(ports)

	// c := &serial.Config{Name: "COM4", Baud: 9600, ReadTimeout: time.Second * 2}
	// port, err := serial.OpenPort(c)
	// if err != nil {
	// 	zlog.Error(err)
	// }
	// s := sensor.NewVoltageSensor(port)

	for _, v := range ports {
		c := &sensor.Config{Name: v, Baud: 9600, ReadTimeout: time.Second * 2}

		port, err := sensor.OpenPort(c)
		if err != nil {
			zlog.Error(err)
			continue
		}
		
		s := sensor.NewVoltageSensor(port)
	}
}