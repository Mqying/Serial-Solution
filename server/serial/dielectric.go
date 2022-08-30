package serial

import (
	"errors"
	"fmt"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/zlog"
	"go.uber.org/zap"
)

var (
	dateTimeForIII string = "" // Date time not same in the same group on 3/6-cups version dielectron device
)

func GetDielectronRecordStruct(b []byte) (*model.DielectronRecord, error) {
	var (
		detectionTime string
		average       float64
		index         int
		items         [10]float64
		err error
		versionIII bool
	)

	itemsLen := len(items)
	itemUnit := 9
	k := 0
	itemIndexStart := 0
	itemIndexEnd := 0

	// get items data
	for i := 0; i < len(b); i++ {
		k = i

		// ASCLL 0X5b = [
		for b[k] != 0x5b {
			k++
		}
		itemIndexStart = k

		// ASCLL x05d = ]
		for b[k] != 0x5d {
			k++
		}
		itemIndexEnd = k

		if !IsNum(b[itemIndexStart+1]) {
			zlog.Debug("[DIELECTRIC] itme index error", zap.ByteString("index: ", b[itemIndexStart+1:itemIndexStart+2]))

			err := errors.New("[DIELECTRIC] itme index error")
			return nil, err
		}

		// get item index
		var itemIndex int
		if itemIndexEnd-itemIndexStart > 2 {
			if !IsNum(b[itemIndexStart+2]) {
				zlog.Debug("[DIELECTRIC] itme index error", zap.ByteString("index: ", b[itemIndexStart+2:itemIndexStart+3]))

				err := errors.New("[DIELECTRIC] itme index error")
				return nil, err
			}

			itemIndex = int(b[itemIndexStart+1]-0x30)*10 + int(b[itemIndexStart+2]-0x30) - 1
		} else {
			itemIndex = int(b[itemIndexStart+1]-0x30) - 1
		}

		// get item value
		itemIntegerOffset := 1
		itemDecimalOffset := 5

		if !IsNum(b[k+itemIntegerOffset]) || !IsNum(b[k+itemIntegerOffset+1]) {
			zlog.Debug("[DIELECTRIC] itme value error", zap.ByteString("get value: ", b[k+itemIntegerOffset:k+itemIntegerOffset+2]))

			err := errors.New("[DIELECTRIC] itme value error")
			return nil, err
		}

		itemInteger := float64(int(b[k+itemIntegerOffset]-0x30)*10 + int(b[k+itemIntegerOffset+1]-0x30))
		itemDecimal := float64(b[k+itemDecimalOffset]-0x30) / 10
		item := itemInteger + itemDecimal

		items[itemIndex] = item

		i = k + itemDecimalOffset

		if itemIndex == itemsLen-1 {
			break
		}
	}
	k += itemUnit + 1 //

	for b[k] != ':' {
		k++
	}
	k++

	// average
	averageDecimalOffset := 4
	if !IsNum(b[k]) || !IsNum(b[k+1]) || !IsNum(b[k+averageDecimalOffset]) {
		zlog.Debug("[DIELECTRIC] format average error", zap.ByteString("average: ", b[k:k+averageDecimalOffset+1]))

		err := errors.New("[DIELECTRIC] format average error")
		return nil, err
	}
	averageInterger := float64(int(b[k]-0x30)*10 + int(b[k+1]-0x30))
	averageDecimal := float64(b[k+averageDecimalOffset]-0x30) / 10
	average = averageInterger + averageDecimal
	k += averageDecimalOffset + 1

	for !IsNum(b[k]) {
		k++
	}

	// idx
	if !IsNum(b[k]) {
		zlog.Debug("[DIELECTRIC] format idx error", zap.Binary("idx: ", b[k:k+1]))

		err := errors.New("[DIELECTRIC] format idx error")
		return nil, err
	} else {
		if IsNum(b[k+1]) {
			index = int(b[k] - 0x30) * 10 + int(b[k+1] - 0x30)
			k += 2
		} else {
			versionIII = true // on 3/6-cups version, index is one byte length

			index = int(b[k] - 0x30)
			k+= 1
		}
	}

	for !IsNum(b[k]) {
		k++
	}

	// date
	yearOffset := k
	monthOffset := k + 6
	dayOffset := k + 10
	hourOffset := k + 14
	minuteOffset := k + 18
	dateStr := ""

	dateStr += string(b[yearOffset:yearOffset+4]) + "-" +
		string(b[monthOffset:monthOffset+2]) + "-" +
		string(b[dayOffset:dayOffset+2]) + " " +
		string(b[hourOffset:hourOffset+2]) + ":" +
		string(b[minuteOffset:minuteOffset+2]) + ":00"

	if !IsNum(b[yearOffset]) || !IsNum(b[yearOffset + 1]) || !IsNum(b[yearOffset + 2]) || !IsNum(b[yearOffset + 3]) {
		zlog.Debug("[DIELECTRIC] format time-year error", zap.ByteString("year: ", b[yearOffset:yearOffset+3]))

		err := errors.New("[DIELECTRIC] format time-year error")
		return nil, err
	}

	if !IsNum(b[monthOffset]) || !IsNum(b[monthOffset + 1]) {
		zlog.Debug("[DIELECTRIC] format time-month error", zap.ByteString("month: ", b[monthOffset:monthOffset+1]))

		err := errors.New("[DIELECTRIC] format time-month error")
		return nil, err
	}

	if !IsNum(b[dayOffset]) || !IsNum(b[dayOffset + 1]) {
		zlog.Debug("[DIELECTRIC] format time-day error", zap.ByteString("day: ", b[dayOffset:dayOffset+1]))

		err := errors.New("[DIELECTRIC] format time-day error")
		return nil, err
	}

	if !IsNum(b[hourOffset]) || !IsNum(b[hourOffset + 1]) {
		zlog.Debug("[DIELECTRIC] format time-hour error", zap.ByteString("hour: ", b[hourOffset:hourOffset+1]))

		err := errors.New("[DIELECTRIC] format time-hour error")
		return nil, err
	}

	if !IsNum(b[minuteOffset]) || !IsNum(b[minuteOffset + 1]) {
		zlog.Debug("[DIELECTRIC] format time-min error", zap.ByteString("min: ", b[minuteOffset:minuteOffset+1]))

		err := errors.New("[DIELECTRIC] format time-min error")
		return nil, err
	}

	year := int(b[yearOffset] - 0x30) * 1000 + int(b[yearOffset + 1] - 0x30) * 100 + int(b[yearOffset + 2] - 0x30) * 10 + int(b[yearOffset + 3] - 0x30)
	mon := int(b[monthOffset] - 0x30) * 10 + int(b[monthOffset + 1] - 0x30)
	day := int(b[dayOffset] - 0x30) * 10 + int(b[dayOffset + 1] - 0x30)
	hour := int(b[hourOffset] - 0x30) * 10 + int(b[hourOffset + 1] - 0x30)
	min := int(b[minuteOffset] - 0x30) * 10 + int(b[minuteOffset + 1] - 0x30)

	if detectionTime, err = getDielectronTime(year, mon, day, hour, min); err != nil {
		zlog.Error(err)
		return nil, err
	} else {
		if versionIII {
			if index == 1 { // 3/6-cups version should guarantee the following seq: 1 -> 2 -> 3 -> ...
				dateTimeForIII = detectionTime
			} else {
				if dateTimeForIII != "" { // if sequence not correct, show datetime from data packet
					detectionTime = dateTimeForIII
				}
			}
		}

		return &model.DielectronRecord{
			DetectionTime: detectionTime,
			Average:       average,
			Index:         index,
			Items:         items,
		}, nil
	}
}

func getDielectronTime(year, mon, day, hour, minute int) (string, error) {
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
	// t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)

	return timeStr, nil
}