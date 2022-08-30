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

		fmt.Println("Option = ", flag)
	}

	if flag == cmd.SerialCmdEmpty {
		data := []byte{
			29,100, 1, 27, 64, 91, 54, 93, 32, 80, 72,
			32, 73, 46, 53, 13, 91,53, 93, 32, 80, 72, 
			32, 49, 46, 53, 13,91, 52, 93, 32, 80, 72, 
			32, 50, 46, 53, 13, 91, 51, 93, 32, 80, 
			72, 32, 73, 46,53, 13, 91, 50, 93, 
			32, 80, 72, 32, 73, 46, 53,
			13, 91, 49, 93, 32, 80, 72, 32, 73, 46,
			53, 13, 178, 226, 202, 212, 202, 177, 188, 
			228, 58, 49, 53, 212, 194, 73, 53, 200, 213, 
			49, 53, 202, 177, 73, 53, 183, 214, 13, 177, 
			224, 186, 197, 58, 49, 13, 27, 74, 80, 13, 29, 100, 1,
		}

		if s, err := serial.GetAcidRecordStruct(data); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%v\n", s)
		}

		return
	}

	deviceType := model.AcidType

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

	if s, err := serial.GetAcidRecordStruct(data); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%v\n", s)
	}
}