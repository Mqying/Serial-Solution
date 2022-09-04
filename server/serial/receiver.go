package serial

import (
	"errors"

	"github.com/silverswords/dielectric/server/zlog"
)

var (
	errDeviceNoResp = errors.New("remote device no response")
)

// To get all data from serial, there are three ways
// I.   Set max time
// II.  Fixed length
// III. Analyse end data, like ReceiveDielectronData below

func (s *VoltageSensor) ReceiveDielectronData() ([]byte, error) {
	data := []byte{}
	l := 0

	empty := 0
	for {
		buffer := make([]byte, 10)

		n, err := s.Port.Read(buffer)
		if err != nil {
			zlog.Error(err)
			return nil, err
		}

		if n == 0 {
			zlog.Debug("[Dielectron] Received zero byte")

			empty += 1
			if empty == 3 {
				break
			}
		}

		data = append(data, buffer[:n]...)

		l = len(data)
		if (data[l-1] == 0xD6 && data[l-2] == 0xB7) ||
			(data[l-1] == 0x20 && data[l-2] == 0xD6) {
			info := LogBytes(data)
			zlog.Debug("[Dielectron] ALL bytes pc get : " + info)

			return data, nil
		}
	}

	if len(data) == 0 {
		return nil, errDeviceNoResp
	}

	return data, nil
}

func (s *VoltageSensor) ReceiveAcidData() ([]byte, error) {

	data := []byte{}
	var length int

	for {
		buffer := make([]byte, 10)

		n, err := s.Port.Read(buffer)
		if err != nil {
			zlog.Error(err)
			return nil, err
		}

		data = append(data, buffer[:n]...)
		length = len(data)

		if (length > 3) && (data[length-1] == 0x01 && data[length-2] == 0x64 && data[length-3] == 0x1d) {
			info := LogBytes(data)
			zlog.Debug("[Acid] ALL bytes pc get : " + info)

			return data, nil
		}
	}
}

func (s *VoltageSensor) ReceiveFlashData() ([]byte, error) {
	data := []byte{}
	length := 0

	retry := 0
	for {
		if retry == 5 {
			break
		}

		buffer := make([]byte, 10)

		n, err := s.Port.Read(buffer)
		if err != nil {
			zlog.Error(err)
			return nil, err
		}

		if n == 0 {
			retry++
			continue
		} else {
			retry = 0
		}

		data = append(data, buffer[:n]...)
		length = len(data)

		if (data[length-1] == 13) && !IsAscii(data[length-2]) {
			info := LogBytes(data)
			zlog.Debug("[Flash] ALL bytes pc get : " + info)

			return data, nil
		}
	}

	return data, nil
}

func (s *VoltageSensor) ReceiveWaterData() ([]byte, error) {
	data := []byte{}
	length := 0

	for {
		buffer := make([]byte, 20)

		n, err := s.Port.Read(buffer)
		if err != nil {
			zlog.Error(err)

			return nil, err
		}

		data = append(data, buffer[:n]...)

		length = len(data)
		if (length > 3) && (data[length-1] == 0x01 && data[length-2] == 0x64 && data[length-3] == 0x1d) {
			zlog.Debug("[Water] ALL bytes pc get : " + LogBytes(data))
			return data, nil
		}
	}
}

func (s *VoltageSensor) ReceiveWaterUselessData() error {
	data := []byte{}

	for {
		buffer := make([]byte, 10)

		n, err := s.Port.Read(buffer)
		if err != nil {
			zlog.Error(err)

			return err
		}

		data = append(data, buffer[:n]...)
		if len(data) >= 8 {
			return nil
		}
	}
}
