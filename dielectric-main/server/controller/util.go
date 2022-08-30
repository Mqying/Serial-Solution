package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	data "github.com/silverswords/dielectric/server/model"
	recordmodel "github.com/silverswords/dielectric/server/serial"
	"github.com/silverswords/dielectric/server/zlog"
)

//parse makes data transform from []byte to *model.Record
func (c *DeviceController) Parse(gbkData []byte, deviceType int) (interface{}, error) {
	switch deviceType {
	case data.DielectronType:
		record, err := recordmodel.GetDielectronRecordStruct(gbkData) // change GetInfoByStr function name
		if err != nil {
			return nil, err
		}

		if err := data.CheckAndInsertDielectronRecord(c.db, record); err != nil {
			return nil, err
		}

		return record, nil

	case data.WaterType:
		record, err := recordmodel.GetWaterRecordStruct(gbkData) // change GetInfoByStr function name
		if err != nil {
			return nil, err
		}

		if err := data.CheckAndInsertWaterRecord(c.db, record); err != nil {
			return nil, err
		}

		return record, nil

	case data.AcidType:
		record, err := recordmodel.GetAcidRecordStruct(gbkData) // change GetInfoByStr function name
		if err != nil {
			return nil, err
		}

		if err := data.CheckAndInsertAcidRecord(c.db, record); err != nil {
			return nil, err
		}

		return record, nil

	case data.FlashType:
		record, err := recordmodel.GetFlashRecordStruct(gbkData) // change GetInfoByStr function name
		if err != nil {
			return nil, err
		}

		if err := data.CheckAndInsertFlashRecord(c.db, record); err != nil {
			return nil, err
		}

		return record, nil
	}

	err := fmt.Errorf("wrong device type")
	return nil, err
}

func GetDeviceType(ctx *gin.Context) (int, error) {
	value := ctx.Query("type")
	if value == "" {		
		err := errors.New("wrong device type")
		zlog.Error(err)
		return -1, err
	}

	deviceType, err := strconv.Atoi(value)
	if err != nil {
		zlog.Error(err)
		return -1, err
	}

	return deviceType, err
}
