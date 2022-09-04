package serial

import (
	"errors"
	"strconv"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/zlog"
)

// To simplify zlog usage and seperate different machine

// To seperate machines
func (s *VoltageSensor) ReceiveData(deviceType int) ([]byte, error) {
	switch deviceType {
	case model.DielectronType:
		data, err := s.ReceiveDielectronData()
		if err != nil {
			return nil, err
		}

		return data, nil

	case model.WaterType:
		data, err := s.ReceiveWaterData()
		if err != nil {
			return nil, err
		}
		return data, nil

	case model.AcidType:
		data, err := s.ReceiveAcidData()
		if err != nil {
			return nil, err
		}
		return data, nil

	case model.FlashType:
		data, err := s.ReceiveFlashData()
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	err := errors.New("wrong device type")
	zlog.Error(err)
	return nil, err
}

// To help analyse data
func IsAscii(b byte) bool {
	return b <= 127
}

func IsNum(b byte) bool {
	if b >= 48 && b <= 57 {

		return true
	}

	return false
}

func IsAlpha(b byte) bool {
	if (b >= 65 && b <= 90) || (b >= 97 && b <= 122) {

		return true
	}

	return false
}

// To simplify zlog usage
func LogBytes(b []byte) string {
	if len(b) == 0 {
		err := errors.New("-----------NO DATA------------ ")
		zlog.Error(err)

		return "----NO DATA ERR----"
	}

	if len(b) == 1 {
		return strconv.Itoa(int(b[0]))
	}

	info := strconv.Itoa(int(b[0]))

	for _, v := range b[1:] {
		info = info + ", " + strconv.Itoa(int(v))
	}

	length := len(b)
	info = info + " ------- total len : " + strconv.Itoa(length)

	return info
}

func LogInts(b []int) string {
	if len(b) == 0 {
		err := errors.New("-----------NO DATA------------ ")
		zlog.Error(err)

		return "----NO DATA ERR----"
	}

	if len(b) == 1 {
		return strconv.Itoa(b[0])
	}

	info := strconv.Itoa(b[0])

	for _, v := range b[1:] {
		info = info + ", " + strconv.Itoa(v)
	}

	length := len(b)
	info = info + " ------- total len : " + strconv.Itoa(length)

	return info
}

func LogStrs(b []string) string {
	if len(b) == 0 {
		err := errors.New("-----------NO DATA------------ ")
		zlog.Error(err)

		return "----NO DATA ERR----"
	}

	if len(b) == 1 {
		return b[0]
	}

	info := b[0]

	for _, v := range b[1:] {
		info = info + ", " + v
	}

	length := len(b)
	info = info + " ------- total len : " + strconv.Itoa(length)

	return info
}
