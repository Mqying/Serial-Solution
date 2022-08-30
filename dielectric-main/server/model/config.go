package model

import (
	"database/sql"
	"errors"
	"fmt"
)

const (
	DeviceTypeTableName = "device_type"
	LogoTableName       = "logo"
)

const (
	configSqliteDeviceTypeGet = iota
	configSqliteLogoGet
)

var (
	configSQLString = []string{
		fmt.Sprintf(`SELECT water, dielectron, falsh, acid FROM %s LIMIT 1`, DeviceTypeTableName),
		fmt.Sprintf(`SELECT enable FROM %s LIMIT 1`, LogoTableName),
	}
)

type Config struct {
	Dielectron int `json:"dielectron"`
	Water      int `json:"water"`
	Acid       int `json:"acid"`
	Flash      int `json:"flash"`
	Logo       int `json:"logo"`
}

func GetConfig(db *sql.DB) (*Config, error) {
	var (
		dielectron int
		water      int
		acid       int
		flash      int
		logo       int
	)

	err := db.QueryRow(configSQLString[configSqliteDeviceTypeGet]).Scan(&water, &dielectron, &flash, &acid)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(configSQLString[configSqliteLogoGet]).Scan(&logo)
	if err != nil {
		return nil, err
	}

	config := &Config{
		Dielectron: dielectron,
		Water:      water,
		Acid:       acid,
		Flash:      flash,
		Logo:       logo,
	}
	
	return config, nil
}

func GetDeviceType(db *sql.DB) (int, error) {
	var (
		dielectron int
		water      int
		acid       int
		flash      int
	)

	err := db.QueryRow(configSQLString[configSqliteDeviceTypeGet]).Scan(&water, &dielectron, &flash, &acid)
	if err != nil {
		return -1, err
	}

	if dielectron == 1 {
		return DielectronType, nil 
	}

	if water == 1 {
		return WaterType, nil
	}

	if acid == 1 {
		return AcidType, nil
	}

	if flash == 1 {
		return FlashType, nil
	}

	err = errors.New("no device type")

	return -1, err
}