package serial

import (
	"errors"
	"fmt"
	"time"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/zlog"
	"go.uber.org/zap"
)

//大气压强:321.0 预闪温度:256 闪点温度:123.0 样品编号:654321 测试时间:07月03日11时52分
func GetFlashRecordStruct(b []byte) (*model.FlashRecord, error) {
	var (
		id			   	int
		preTemp      	float64  		//预闪温度
		pointTemp       float64  		//闪点温度
		pressure        float64 	
		time			time.Time
		qian			int
		bai				int
		shi				int
		ge				int
		mon				int
		day				int
		hour			int
		min  			int

		errorNum		int
	)

	idx := 0

	for IsAscii(b[idx]) {
		idx++
	}

	for !IsAscii(b[idx]) { // 跳过 大气压强:
		idx++
	}

	if b[idx] == ':' {
		idx++
	}

	// get pressure
	if !IsNum(b[idx]) || !IsNum(b[idx + 1]) || !IsNum(b[idx + 2]) || !IsNum(b[idx + 4]) {
		errorNum += 1
		pressure = 0

		zlog.Debug("[Flash] Wrong pressure data:", zap.ByteString("pressure", b[idx:idx + 5]))
	} else {
		qian = int(b[idx]     - 0x30)
		bai  = int(b[idx + 1] - 0x30)
		shi	 = int(b[idx + 2] - 0x30)
		ge   = int(b[idx + 4] - 0x30)

		pressure = float64(qian) * 100 + float64(bai) * 10 + float64(shi) + float64(ge) * 0.1
	}
	idx += 6

	// get pre temperature
	for !IsAscii(b[idx]) { // 跳过 预闪温度
		idx++
	}

	if b[idx] == ':' {
		idx++
	}

	if !IsNum(b[idx]) || !IsNum(b[idx + 1]) || !IsNum(b[idx + 2]) {
		errorNum += 1
		zlog.Debug("[Flash] Wrong previous temperature data:", zap.ByteString("pressure", b[idx:idx + 3]))
		preTemp = 0
	} else {
		bai = int(b[idx]	 - 	0x30)
		shi = int(b[idx + 1] - 	0x30)
		ge  = int(b[idx + 2] - 	0x30)

		preTemp = float64(bai) * 100 + float64(shi) * 10 + float64(ge)
	}
	idx += 4

	for !IsAscii(b[idx]) { // 跳过 闪点温度
		idx++
	}

	if b[idx] == ':' {
		idx++
	}

	if !IsNum(b[idx]) || !IsNum(b[idx + 1]) || !IsNum(b[idx + 2]) {
		errorNum += 1
		zlog.Debug("[Flash] Wrong point temperature data:", zap.ByteString("point temperature", b[idx:idx + 5]))
		pointTemp = 0
	} else {
		qian = int(b[idx]     - 0x30)
		bai  = int(b[idx + 1] - 0x30)
		shi	 = int(b[idx + 2] - 0x30)
		ge   = int(b[idx + 4] - 0x30)

		pointTemp = float64(qian) * 100 + float64(bai) * 10 + float64(shi) + float64(ge) * 0.1
	}
	idx += 6

	if b[idx] == 0xd {
		idx++
	}

	for !IsAscii(b[idx]) { // 跳过 样品编号
		idx++
	} 

	if b[idx] == ':' {
		idx++
	}

	var index  [6]int
	var idFlag bool				

	for i := 0; i < 6; i++ {
		if !IsNum(b[idx + i]) {
			errorNum++
			zlog.Debug("[Flash] Wrong index data:", zap.ByteString("index data", b[idx:idx + 6]))
			idFlag = true
		} else {
			index[i] = int(b[idx + i] - 0x30)
		}
	}

	if !idFlag {
		id = index[5] + index[4] * 10 + index[3] * 100 + index[2] * 1000 + 
			 index[1] * 10000 + index[0] * 100000
	} else {
		id = 0
	}
	idx += 7

	//get mon
	for !IsAscii(b[idx]) { // 跳过 测试时间
		idx++
	}

	if b[idx] == ':' {
		idx++
	}

	if !IsNum(b[idx]) || !IsNum(b[idx + 1]) {
		zlog.Debug("[Flash] Wrong mon data:", zap.ByteString("mon data", b[idx:idx + 2]))
		mon = 0
	} else {
		shi = int(b[idx]     - 0x30)
		ge  = int(b[idx + 1] - 0x30)
		mon = int(shi) * 10 + int(ge)
	}
	idx += 2

	//get day
	for !IsAscii(b[idx]) { // 跳过 月
		idx++
	}

	if !IsNum(b[idx]) || !IsNum(b[idx + 1]) {
		zlog.Debug("[Flash] Wrong day data:", zap.ByteString("day data", b[idx:idx + 2]))
		day = 0
	} else {
		shi = int(b[idx] 	 - 0x30)
		ge  = int(b[idx + 1] - 0x30)
		day = int(shi) * 10 + int(ge)
	}
	idx += 2

	//get hour
	for !IsNum(b[idx]) { // 跳过  日
		idx++
	}

	if !IsNum(b[idx]) || !IsNum(b[idx + 1]) {
		zlog.Debug("[Flash] Wrong hour data:", zap.ByteString("mon hour", b[idx:idx + 2]))
		hour = 0 
	} else {
		shi = int(b[idx] 	 - 0x30)
		ge  = int(b[idx + 1] - 0x30)
		hour = int(shi) * 10 + int(ge)
	}
	idx += 2

	//get min
	for !IsNum(b[idx]) { // 跳过  时
		idx++
	}

	if !IsNum(b[idx]) || !IsNum(b[idx + 1]) {
		zlog.Debug("[Flash] Wrong min data:", zap.ByteString("min data", b[idx:idx + 2]))
		min = 0 
	} else {
		shi = int(b[idx] 	 - 0x30)
		ge  = int(b[idx + 1] - 0x30)
		min = int(shi) * 10 + int(ge)
	}

	time, err := getFlashTime(22 ,mon, day, hour, min)

	if (errorNum == 0) && (err == nil) {
		return &model.FlashRecord{
			DetectionTime:	time,
			Id:				id,
			Pretemp: 		preTemp,
			Pressure: 		pressure,
			Pointtemp: 		pointTemp,
		}, nil 
	} else {
		err := errors.New("[FLASH] Wrong flash data")
		return nil, err 
	}
}

func getFlashTime(year, mon, day, hour, minute int) (time.Time, error) {
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