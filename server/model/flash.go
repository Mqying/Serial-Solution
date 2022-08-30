package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/silverswords/dielectric/server/zlog"
)

const FlashTableName = "Flash"

const (
	FlashSqliteRecordCreateTable = iota
	FlashSqliteRecordInsert
	FlashSqliteRecordGetAll
	FlashSqliteRecordIsExist
	FlashSqliteGetRecordByid
)

var (
	FlashRecordSQLString = []string{
		fmt.Sprintf(`CREATE  TABLE IF NOT EXISTS %s (
			detection_time   TIMESTAMP NOT NULL PRIMARY KEY,
			id			 	 INTEGER  NOT NULL DEFAULT 0,
			pretemp 		 REAL NOT NULL DEFAULT 0,
			pointtemp		 REAL NOT NULL DEFAULT 0,
			pressure 		 REAL NOT NULL DEFAULT 0,
			create_time 	 TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`, FlashTableName),
		fmt.Sprintf(`INSERT INTO %s(detection_time, id, pretemp, pointtemp, pressure) VALUES (?, ?, ?, ?, ?);`, FlashTableName),
		fmt.Sprintf(`SELECT detection_time, id, pretemp, pointtemp, pressure FROM %s ORDER BY detection_time DESC`, FlashTableName),
		fmt.Sprintf(`SELECT COUNT(1) FROM %s WHERE detection_time = ?`, FlashTableName),
		fmt.Sprintf(`SELECT detection_time, id, pretemp, pointtemp, pressure FROM %s WHERE id = ?`, FlashTableName),
	}
)

type FlashRecord struct {
	DetectionTime 	time.Time `json:"detection_time,omitempty"`
	Id         		int       `json:"sample_number,omitempty"`
	Pretemp      	float64   `json:"pre_temperature,omitempty"`
	Pointtemp       float64   `json:"point_temperature,omitempty"`
	Pressure        float64   `json:"pressure,omitempty"`
}

func CreateFlashRecordTable(db *sql.DB) error {
	_, err := db.Exec(FlashRecordSQLString[FlashSqliteRecordCreateTable])
	return err
}

func CheckAndInsertFlashRecord(db *sql.DB, d *FlashRecord) error {
	isExist, err := FlashRecordIsExist(db, d.DetectionTime)
	if err != nil {
		return err
	}
	if isExist {
		return nil 
	}

	result, err := db.Exec(FlashRecordSQLString[FlashSqliteRecordInsert],
		d.DetectionTime, d.Id, d.Pretemp, d.Pointtemp, d.Pressure)
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

func GetAllFlashRecords(db *sql.DB) ([]*FlashRecord, error) {
	var (
		detectionTime  	time.Time
		id		   		int
		pretemp      	float64  
		pointtemp       float64  	
		pressure        float64 	

		result []*FlashRecord
	)

	rows, err := db.Query(FlashRecordSQLString[FlashSqliteRecordGetAll])
	if err != nil {
		zlog.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&detectionTime, &id, &pretemp, &pointtemp, &pressure); err != nil {
			zlog.Error(err)
			return nil, err
		}
		data := &FlashRecord{
			DetectionTime: 	detectionTime,
			Id: 			id,
			Pretemp: 		pretemp,
			Pointtemp: 		pointtemp,
			Pressure: 		pressure,
		}
		result = append(result, data)
	}

	return result, nil
}

func FlashRecordIsExist(db *sql.DB, detectionTime time.Time) (bool, error) {
	row := db.QueryRow(FlashRecordSQLString[FlashSqliteRecordIsExist], detectionTime)
	if err := row.Err(); err != nil {
		zlog.Error(err)
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