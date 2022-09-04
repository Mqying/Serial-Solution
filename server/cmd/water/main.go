package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/silverswords/dielectric/server/cmd"
	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/serial"
)

func main() {
	var flag int
	var err error

	if len(os.Args) == 2 {
		flag, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(flag)
	}

	if flag == cmd.SerialCmdEmpty {
		data := []byte{
			29, 100, 1, 27, 64, 177, 224, 186, 197, 58, 49, 13, 186,
			172, 203, 174, 193, 191, 58, 50, 49, 55, 46, 50, 117, 103,
			13, 178, 226, 202, 212, 202, 177, 188, 228, 58,

			50, 50, // 22
			46,
			48, 56, // 08
			46,
			48, 56, // 08
			46, 32,
			49, 52, // 14
			58,
			50, 56, // 28

			13, 27, 74, 80, 13, 29, 100, 1,
		}

		if s, err := serial.GetWaterRecordStruct(data); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%v\n", s)
		}

		return
	}

	deviceType := model.WaterType

	v := &serial.VoltageSensor{}

	var data []byte

	if err = v.OpenSerial(deviceType); err != nil {
		fmt.Println("Open Serial Error:", err)
		return
	}

	if flag == cmd.SerialCmdFrontPage {
		if data, err = v.SendFrontPage(deviceType); err != nil {
			fmt.Println(flag, " Error ", err)
			return
		}
	}

	if flag == cmd.SerialCmdPreviousPage {
		if data, err = v.SendPreviousPage(deviceType); err != nil {
			fmt.Println(flag, " Error ", err)
			return
		}
	}

	if flag == cmd.SerialCmdNextPage {
		if data, err = v.SendNextPage(deviceType); err != nil {
			fmt.Println(flag, " Error ", err)
			return
		}
	}

	if flag == cmd.SerialCmdPrint {
		v.SendPrint()
		return
	}

	if s, err := serial.GetWaterRecordStruct(data); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%v\n", s)
	}
}
