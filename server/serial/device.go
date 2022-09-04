package serial

import (
	"bufio"
	"errors"
	"time"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/zlog"
	"github.com/tarm/serial"
	stserial "go.bug.st/serial"
	"go.uber.org/zap"
)

// Set all commands; Req -> Controller -> Device -> Machine -> Parse -> Models -> DB -> Resp

var (
	frontPage    = []byte{0x5A, 0xAA, 0xA5}
	nextPage     = []byte{0x5A, 0xBB, 0xA5}
	previousPage = []byte{0x5A, 0xCC, 0xA5}
	print        = []byte{0x5A, 0xDD, 0xA5}
	getId		 = []byte{0x5A, 0xEE, 0xA5}

	waterGetId	  	  = []byte{0x5A, 0xAA, 0xA5}
	waterFrontPage	  = []byte{0x5A, 0xBB, 0xA5}
	waterNextPage 	  = []byte{0x5A, 0xCC, 0xA5}
	waterPreviousPage = []byte{0x5A, 0xDD, 0xA5}
	waterPrint		  = []byte{0x5A, 0xEE, 0xA5}

	ID []byte    = []byte{90, 92}
	dt int		 //get deviceType
)

type VoltageSensor struct {
	*serial.Port
	*bufio.Reader
}

func NewVoltageSensor(s *serial.Port) *VoltageSensor {
	sensor := &VoltageSensor{
		s,
		bufio.NewReader(s),
	}
	return sensor
}

func (s *VoltageSensor) OpenSerial(deviceType int) error {
	dt = deviceType
	if s.Port != nil {
		s.Port.Close()
	}

	baud := 19200

	if deviceType == model.FlashType {
		baud = 9600
	}

	portList, err := stserial.GetPortsList()
	if err != nil {
		zlog.Info(err.Error())

		return err
	}

	if len(portList) == 0 {
		zlog.Warn("[Serial] No available com device.")

		return errors.New("no serial devices")
	}

	zlog.Info("[Serial] Port List:", zap.Strings("com-interface", portList))

	for _, device := range portList {
		zlog.Info("[Serial] Trying port ", zap.String("open", device))

		config := &serial.Config{
			Name: device, 
			Baud: baud, 
			ReadTimeout: time.Millisecond * 1000,
		}

		if port, err := serial.OpenPort(config); err != nil {
			zlog.Error(err, zap.String("interface-open", device))
			continue
		} else {
			s.Port = port
			s.Reader = bufio.NewReader(port)

			if ID, err = s.GetDeviceID(deviceType); err != nil {
				s.Port.Close()
				s.Port = nil
				s.Reader = nil

				zlog.Error(err, zap.String("interface-read", device))

				continue
			} else {
				zlog.Info("Interface:", zap.String("interface-read", device))
				return nil
			}
		}
	}

	return errors.New("[DEVICE] serial device not attached")
}

func (s *VoltageSensor) SendFrontPage(deviceType int) ([]byte, error) {
	if s.Port == nil {
		if err := s.OpenSerial(deviceType); err != nil {
			return nil, err
		}
	}

	var (
		n	int
		err error
	)

	if deviceType == model.WaterType {
		n, err = s.Write(waterFrontPage)
	} else {
		n, err = s.Write(frontPage)
	}

	if err != nil {
		zlog.Error(err)
		return nil, err
	}

	if n == len(frontPage) {
		zlog.Info("[DEVICE] send frontPage successfully")
	}

	return s.ReceiveData(deviceType)
}

func (s *VoltageSensor) SendNextPage(deviceType int) ([]byte, error) {
	if s.Port == nil {
		if err := s.OpenSerial(deviceType); err != nil {
			return nil, err
		}
	}

	var (
		n   int
		err error
	)

	if deviceType == model.WaterType {
		n, err = s.Write(waterNextPage)
	} else {
		n, err = s.Write(nextPage)
	}

	if err != nil {
		zlog.Error(err)
		return nil, err
	}

	if n == len(nextPage) {
		zlog.Info("[DEVICE] send nextPage successfully")
	}

	return s.ReceiveData(deviceType)
}

func (s *VoltageSensor) SendPreviousPage(deviceType int) ([]byte, error) {
	if s.Port == nil {
		if err := s.OpenSerial(deviceType); err != nil {
			return nil, err
		}
	}

	var (
		n   int 
		err error
	)

	if deviceType == model.WaterType {
		n, err = s.Write(waterPreviousPage)
	} else {
		n, err = s.Write(previousPage)
	}

	if err != nil {
		zlog.Error(err)
		return nil, err
	}

	if n == len(previousPage) {
		zlog.Info("[DEVICE] send previousPage successfully")
	}

	return s.ReceiveData(deviceType)
}

func (s *VoltageSensor) SendPrint() error {
	if s.Port == nil {
		if err := s.OpenSerial(dt); err != nil {
			return err
		}
	}

	var (
		n	int
		err	error
	)

	if dt == model.WaterType {
		n, err = s.Write(waterPrint)
	} else {
		n, err = s.Write(print)
	}

	if err != nil {
		zlog.Error(err)
		return err
	}

	if n == len(print) {
		zlog.Info("[DEVICE] send print successfully")
	}

	return nil
}