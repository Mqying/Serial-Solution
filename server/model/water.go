package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/silverswords/dielectric/server/zlog"
)

const WaterTableName = "Water"

const (
	waterSqliteRecordCreateTable = iota
	waterSqliteRecordInsert
	waterSqliteRecordGetAll
	waterSqliteRecordIsExist
)

var (
	waterRecordSQLString = []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			detection_time  TIMESTAMP NOT NULL PRIMARY KEY,
			quantity   	REAL NOT NULL DEFAULT 0,
			ratio1 		REAL NOT NULL DEFAULT 0,
			ratio2 		REAL NOT NULL DEFAULT 0,
			create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`, WaterTableName),
		fmt.Sprintf(`INSERT INTO %s(detection_time, quantity, ratio1, ratio2) VALUES (?, ?, ?, ?);`, WaterTableName),
		fmt.Sprintf(`SELECT detection_time, quantity, ratio1, ratio2 FROM %s ORDER BY detection_time DESC`, WaterTableName),
		fmt.Sprintf(`SELECT COUNT(1) FROM %s WHERE detection_time = ?`, WaterTableName),
	}
)

type WaterRecord struct {
	DetectionTime  time.Time  `json:"detection_time,omitempty"`
	Index          int        `json:"index,omitempty"`
	Quantity       float64    `json:"quantity,omitempty"`
	Ratio1         float64    `json:"ratio1,omitempty"`
	Ratio2         float64    `json:"ratio2,omitempty"`
}

//CreateTable
func CreateWaterRecordTable(db *sql.DB) error {
	_, err := db.Exec(waterRecordSQLString[waterSqliteRecordCreateTable])
	return err
}

func CheckAndInsertWaterRecord(db *sql.DB, d *WaterRecord) error {
	isExist, err := WaterRecordIsExist(db, d.DetectionTime)
	if err != nil {
		return err
	}
	if isExist {
		return nil 
	}
	
	result, err := db.Exec(waterRecordSQLString[waterSqliteRecordInsert],
		d.DetectionTime, d.Quantity, d.Ratio1, d.Ratio2)
	if err != nil {
		zlog.Error(err)
		return err
	}

	if rows, err := result.RowsAffected(); rows == 0 {
		return err
	}

	if _, err = result.LastInsertId(); err != nil {
		zlog.Error(err)
		return err
	}

	return nil
}

func GetAllWaterRecords(db *sql.DB) ([]*WaterRecord, error) {
	var (
		detectionTime  time.Time
		quantity       float64
		ratio1         float64
		ratio2 		   float64

		result []*WaterRecord
	)

	rows, err := db.Query(waterRecordSQLString[waterSqliteRecordGetAll])
	if err != nil {
		zlog.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&detectionTime, &quantity, &ratio1, &ratio2); err != nil {
			zlog.Error(err)
			return nil, err
		}
		data := &WaterRecord{
			DetectionTime: 	detectionTime,
			Quantity:      	quantity,
			Ratio1:         ratio1,
			Ratio2: 		ratio2,	
		}
		result = append(result, data)
	}

	return result, nil
}

func WaterRecordIsExist(db *sql.DB, detectionTime time.Time) (bool, error) {
	row := db.QueryRow(waterRecordSQLString[waterSqliteRecordIsExist], detectionTime)
	if err := row.Err(); err != nil {
		return false, err
	}

	count := 0
	if err := row.Scan(&count); err != nil {
		zlog.Error(err)
		return false, err
	}

	if count == 1 {
		return true, nil
	}

	return false, nil
}

