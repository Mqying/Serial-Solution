package serial

import (
	"fmt"
	"time"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/zlog"
	"go.uber.org/zap"
)

//[6] PH 9.6 [5] PH 9.5 [4] PH 9.4 [3] PH 9.3 [2] PH 9.2 [1] PH 9.1 测试时间:12月22日08时15分 编号:8
func GetAcidRecordStruct(data []byte) (*model.AcidRecord, error) {
	var (
		month         int
		day           int
		hour          int
		min           int
		no            int
		date_time     [4]*int
		date_time_msg [4]string
	)
	getYear := time.Now().Year()
	data_len := len(data)
	ph := [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
	ph_len := len(ph)
	date_time = [4]*int{&month, &day, &hour, &min}
	date_time_msg = [4]string{
		"month",
		"day",
		"hour",
		"minute",
	}

	const ph_unit = 11
	k := 5
	errorsNumNum := 0

	for i := 0; i < ph_len; i++ {
		idx := ph_len - i - 1

		PH_OFFSET_NO := 1
		if int(data[k+i*ph_unit+PH_OFFSET_NO]-0x30) != (idx + 1) {
			errorsNumNum++
			zlog.Debug("[ACID] Wrong ph info", zap.ByteString("PH", data[k+i*ph_unit+PH_OFFSET_NO:k+i*ph_unit+PH_OFFSET_NO+1]))
			goto finish
		}

		PH_OFFSET_CHAR_P := 4
		if data[k+i*ph_unit+PH_OFFSET_CHAR_P] != 'P' ||
			data[k+i*ph_unit+PH_OFFSET_CHAR_P+1] != 'H' {
			errorsNumNum++
			zlog.Debug("[ACID] Wrong PH bytes", zap.ByteString("PH", data[k+i*ph_unit+PH_OFFSET_CHAR_P:k+i*ph_unit+PH_OFFSET_CHAR_P+2]))
			goto finish
		}

		PH_OFFSET_INT_PLACE := 7
		if !IsNum(data[k+i*ph_unit+PH_OFFSET_INT_PLACE]) {
			zlog.Debug("[ACID] data isn't a number", zap.ByteString("data", data[k+i*ph_unit+PH_OFFSET_INT_PLACE:k+i*ph_unit+PH_OFFSET_INT_PLACE+1]))
			errorsNumNum++

			continue
		}

		PH_OFFSET_DECIMAL_PLACE := 9
		if !IsNum(data[k+i*ph_unit+PH_OFFSET_DECIMAL_PLACE]) {
			zlog.Debug("[ACID] data isn't a number", zap.ByteString("data", data[k+i*ph_unit+PH_OFFSET_DECIMAL_PLACE:k+i*ph_unit+PH_OFFSET_DECIMAL_PLACE+1]))
			errorsNumNum++

			continue
		}

		ph[idx] = float64(data[k+i*ph_unit+PH_OFFSET_INT_PLACE]-0x30)*1.0 + float64(data[k+i*ph_unit+PH_OFFSET_DECIMAL_PLACE]-0x30)*0.1
	}
	k += ph_len * ph_unit

	for i := 0; i < 4; i++ {
		for {
			if !(k < data_len) {
				break
			}
			if !IsAscii(data[k]) {
				k++
			} else {
				break
			}
		}

		if data[k] == ':' {
			k++
		}

		if k == data_len {
			errorsNumNum++
			zlog.Debug("[ACID] missing time")
			goto finish
		}

		if !IsNum(data[k]) {
			errorsNumNum++
			zlog.Debug("[ACID] formate error " + date_time_msg[i] + ":", zap.ByteString(date_time_msg[i], data[k:k+2]))
			*(date_time[i]) = int(data[k+1] - 0x30)
		} else {
			fmt.Printf("%v\n", data[k])
			*(date_time[i]) = int((data[k]-0x30)*10 + (data[k+1] - 0x30))
		}
		k += 2
	}

	for {
		if !(k < data_len) {
			break
		}
		if !IsAscii(data[k]) {
			k++
		} else {
			break
		}
	}

	if data[k] != 0xd {
		errorsNumNum++
		zlog.Debug("[ACID] format error", zap.ByteString("carry", data[k:k+1]))
		goto finish
	}
	k++

	for {
		if !(k < data_len) {
			break
		}
		if !IsAscii(data[k]) {
			k++
		} else {
			break
		}
	}

	if data[k] == ':' {
		k++
	}

	if !IsNum(data[k]) {
		errorsNumNum++
		zlog.Debug("[ACID] format error", zap.ByteString("NO", data[k:k+1]))
		goto finish
	}

	no = int(data[k] - 0x30)

finish:

	yearf := getYear
	monf  := *date_time[0]
	dayf  := *date_time[1]
	hourf := *date_time[2]
	minf  := *date_time[3]

	//time := time.Date(getYear, time.Month(*date_time[0]), int(*(date_time[1])), int(*(date_time[2])), *(date_time[3]), 0, 0, time.Local)

	time, err := getAcidTime(yearf, monf, dayf, hourf, minf)
	if err != nil {
		errorsNumNum++
	}

	if errorsNumNum == 0 {
		return &model.AcidRecord{
			DetectionTime: time,
			Index:         no,
			Items:         ph,
		}, nil
	}

	return &model.AcidRecord{
		DetectionTime: time,
		Index:         no,
		Items:         ph,
	}, err
}

func getAcidTime(year, mon, day, hour, minute int) (time.Time, error) {
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

	timeStr := fmt.Sprintf("%s-%s-%s %s:%s:00", ts[0], ts[1], ts[2], ts[3], ts[4])
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)

	if err != nil {
		return time.Now(), err
	}

	return t, nil
}