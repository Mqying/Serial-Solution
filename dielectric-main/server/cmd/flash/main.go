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
			51, 46, 53, 13, 212, 164, 201, 193, 206, 194, 182, 200, 58, 53, 51, 53,
			13, 201, 193, 181, 227, 206, 194, 182, 200, 58, 113, 53, 51, 46, 53, 13, 
			209, 249, 198, 183, 177, 224, 186, 197, 58, 245, 54, 55, 50, 57, 53, 13, 
			178, 226, 202, 212, 202, 177, 188, 228, 58, 73, 53, 212, 194, 73, 53, 
			200, 213, 32, 73, 53, 202, 177, 73, 53, 183, 214, 13,
		}

		if s, err := serial.GetFlashRecordStruct(data); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%v\n", s)
		}

		return
	}

	deviceType := model.FlashType

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

	if s, err := serial.GetFlashRecordStruct(data); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%v\n", s)
	}
}