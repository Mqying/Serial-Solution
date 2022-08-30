package serial

import (
	"errors"
	"time"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/zlog"
	"go.uber.org/zap"
)

func (s *VoltageSensor) GetDeviceID(deviceType int) ([]byte, error) {
	var (
		n	int 
		err error
	)

	time.Sleep(time.Millisecond * 500)

	if deviceType == model.WaterType {
		n, err = s.Write(waterGetId)
	} else {
		n, err = s.Write(getId)
	}

	if err != nil {
		zlog.Error(err)
		return nil, err
	}

	if n == len(getId) {
		zlog.Info("[ID] send getId successfully")
	}

	data := []byte{}

	retry := 0
	for {
		if retry == 4 {
			break
		}

		buffer := make([]byte, 10)

		n, err := s.Port.Read(buffer)
		if err != nil {
			zlog.Error(err)
			return nil, err
		}

		zlog.Info("[ID] Receiving ID data length: ", zap.Int("id-len", n))

		if n == 0 {
			zlog.Info("[ID] Get ID retry count: ", zap.Int("retry-count", retry))
			retry += 1
			continue
		}

		retry = 0

		zlog.Debug("[ID] Receiving ID data : " + LogBytes(buffer))

		data = append(data, buffer[:n]...)

		if len(data) == 2 {
			zlog.Info("[ID] Get ID SUCCESS")
			return data, nil
		}
	}

	return nil, errors.New("[Serial] device id timeout ")
}