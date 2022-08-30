package model

import (
	"database/sql"
	"errors"

	"github.com/silverswords/dielectric/server/zlog"
)

const (
	DielectronType = iota
	WaterType
	AcidType
	FlashType
)

func CreateRecordTable(db *sql.DB) error {
	err := CreateAcidRecordTable(db)
	if err != nil {
		zlog.Error(err)
		return err
	}

	err = CreateDielectronRecordTable(db)
	if err != nil {
		zlog.Error(err)
		return err
	}

	err = CreateWaterRecordTable(db)
	if err != nil {
		zlog.Error(err)
		return err
	}

	err = CreateFlashRecordTable(db)
	if err != nil {
		zlog.Error(err)
		return err
	}
	
	return nil 
}

func GetAllRecords(db *sql.DB, deviceType int) (interface{}, error) {
	switch deviceType {
	case DielectronType:
		return GetAllDielectronRecords(db)
	case AcidType:
		return GetAllAcidRecords(db)
	case WaterType:
		return GetAllWaterRecords(db)
	case FlashType:
		return GetAllFlashRecords(db)
	}

	return nil, errors.New("wrong device type")
}

