package serial

import (
	"errors"
	"fmt"
	"time"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/zlog"
	"go.uber.org/zap"
)

//txt := "NO:8 MC:999989.7ug Percentage:999989.7PPM TIME: 22.11.11. 10:25"
func GetWaterRecordStruct(b []byte) (*model.WaterRecord, error) {
	var (
		quantity, ratio1, ratio2              float64 = 0.0, 0.0, 0.0
		year, month, day, hour, minute, index int
		error_count                           int = 0
	)

	date_time := [5]*int{&year, &month, &day, &hour, &minute}
	date_msg := [5]string{"year", "month", "day", "hour", "minute"}

	key := 5
	if b[key] == 'N' && b[key+1] == 'O' {
		key += 2
	} else {
		for {
			if IsAscii(b[key]) {
				break
			}
			key++
		}
	}

	if b[key] == ':' {
		key++
	}

	//parse no
	if !IsNum(b[key]) {
		error_count++
		zlog.Debug("[WATER] format err : NO", zap.ByteString("no", b[key:key+1]))
		goto finish
	} else {
		index = int(b[key] - 0x30)
	}
	key += 2

	//pass MC: or 含水量:
	if b[key] == 'M' && b[key+1] == 'C' {
		key += 2
	} else {
		for {
			if IsAscii(b[key]) {
				break
			}
			key++
		}
	}

	if b[key] == ':' {
		key++
	}

	//parse quantity
	for {
		if b[key] == '.' {
			break
		}

		if IsNum(b[key]) {
			quantity = float64(b[key]-0x30) + quantity*10
		} else {
			zlog.Debug("[Water] MC formate:", zap.ByteString("intergral", b[key:key+1]))
		}
		key++
	}
	key++

	//calculate quantity decimals
	if !IsNum(b[key]) {
		zlog.Debug("[Water] MC formate:", zap.ByteString("fractional", b[key:key+1]))
	} else {
		quantity += float64(b[key]-0x30) * 0.1
	}
	key++

	//ug
	if b[key] != 'u' || b[key+1] != 'g' {
		error_count++
		zlog.Debug("[WATER] formate err : ", zap.Any("MC : ", b[key:key+2]))
	}
	key += 3

	//pass Percentage: or 含水率:
	if b[key] == 'P' && b[key+1] == 'e' {
		key += 10
	} else if b[key] == 'T' && b[key+1] == 'I' {
		key += 4
	} else {
		for {
			if IsAscii(b[key]) {
				break
			}
			key++
		}
	}

	if b[key] == ':' {
		key++
	}

	if b[key+9] == ' ' {
		goto parseTime
	}

	//parse Percenrt
	for {
		if !IsNum(b[key]) {
			break
		}

		ratio1 = float64(b[key]-0x30) + ratio1*10
		key++
	}

	if b[key] != '.' {
		error_count++
		zlog.Debug("[WATER] format err percent", zap.ByteString("percent", b[key:]))
	}
	key++

	ratio1 += float64(b[key]-0x30) * 0.1
	key++

	//judge % or PPM
	if b[key] == 'P' && b[key+1] == 'P' && b[key+2] == 'M' {
		key += 3
	} else if b[key] == '%' {
		ratio2 = ratio1
		ratio1 = 0.0

		key += 1
	} else {
		error_count++
		zlog.Debug("[WATER] format err percent", zap.ByteString("unit : ", b[key:key+2]))
	}
	key += 1

	//pass "TIME:" or "13测试时间:"
	if b[key] == 'T' && b[key+1] == 'I' {
		key += 4
	} else {
		for {
			if IsAscii(b[key]) {
				break
			}

			key++
		}
	}

	if b[key] == ':' {
		key++
	}

	//parse time
parseTime:
	for i := 0; i < 3; i++ {
		if !IsNum(b[key]) {
			error_count++
			zlog.Debug(fmt.Sprint("[WATER] format err time : ", date_msg[i]), zap.ByteString("time : ", b[key:key+1]))
			*date_time[i] = int(b[key+1] - 0x30)
		} else {
			*date_time[i] = int(b[key]-0x30)*10 + int(b[key+1]-0x30)
		}
		key += 3
	}

	if b[key] != ' ' {
		error_count++
		zlog.Debug("[WATER] format err time", zap.ByteString("space", b[key:key+1]))
	}
	key++

	if !IsNum(b[key]) {
		error_count++
		zlog.Debug(fmt.Sprint("[WATER] format err time : ", date_msg[3]), zap.ByteString("hour", b[key:key+1]))
		*date_time[3] = int(b[key+1] - 0x30)
	} else {
		*date_time[3] = int(b[key]-0x30)*10 + int(b[key+1]-0x30)
	}
	key += 3

	if !IsNum(b[key]) {
		error_count++
		zlog.Debug(fmt.Sprint("[WATER] format err time : ", date_msg[4]), zap.ByteString("minute", b[key:key+1]))
		*date_time[4] = int(b[key+1] - 0x30)
	} else {
		*date_time[4] = int(b[key]-0x30)*10 + int(b[key+1]-0x30)
	}

finish:
	date, err := getWaterTime(year, month, day, hour, minute)
	if err != nil {
		error_count++
	}

	if error_count == 0 {
		return &model.WaterRecord{
			DetectionTime: date,
			Index:         index,
			Quantity:      quantity,
			Ratio1:        ratio1,
			Ratio2:        ratio2,
		}, nil
	} else {
		err := errors.New("[water] Wrong data")
		return nil, err
	}
}

func getWaterTime(year, mon, day, hour, minute int) (time.Time, error) {
	var ti = make([]int, 5)
	ti = []int{year, mon, day, hour, minute}

	var ts = make([]string, 5)
	for i := 0; i < len(ti); i++ {
		if ti[i] < 10 {
			ts[i] = fmt.Sprintf("0%d", ti[i])
		} else {
			ts[i] = fmt.Sprintf("%d", ti[i])
		}
	}

	timeStr := fmt.Sprintf("20%s-%s-%s %s:%s:00", ts[0], ts[1], ts[2], ts[3], ts[4])
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)

	if err != nil {
		return time.Now(), err
	}

	return t, nil
}