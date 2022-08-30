package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/silverswords/dielectric/server/zlog"
)

const AcidTableName = "Acid"

const (
	acidSqliteRecordCreateTable = iota
	acidSqliteRecordInsert
	acidSqliteRecordGetAll
	acidSqliteRecordIsExist
)

var (
	acidRecordSQLString = []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			detection_time TIMESTAMP NOT NULL PRIMARY KEY,
			item1 REAL NOT NULL DEFAULT 0,
			item2 REAL NOT NULL DEFAULT 0,
			item3 REAL NOT NULL DEFAULT 0,
			item4 REAL NOT NULL DEFAULT 0,
			item5 REAL NOT NULL DEFAULT 0,
			item6 REAL NOT NULL DEFAULT 0,
			create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`, AcidTableName),
		fmt.Sprintf(`INSERT INTO %s(detection_time, item1, item2, item3, item4, item5, item6) VALUES(?, ?, ?, ?, ?, ?, ?);`, AcidTableName),
		fmt.Sprintf(`SELECT detection_time, item1, item2, item3, item4, item5, item6 FROM %s ORDER BY detection_time DESC`, AcidTableName),
		fmt.Sprintf(`SELECT COUNT(1) FROM %s WHERE detection_time = ?`, AcidTableName),
	}
)

type AcidRecord struct {
	DetectionTime time.Time  `json:"detection_time,omitempty"`
	Index         int        `json:"index,omitempty"`
	Items         [6]float64 `json:"items,omitempty"`
}

//CreateTable
func CreateAcidRecordTable(db *sql.DB) error {
	_, err := db.Exec(acidRecordSQLString[acidSqliteRecordCreateTable])
	return err
}

func CheckAndInsertAcidRecord(db *sql.DB, d *AcidRecord) error {
	isExist, err := AcidRecordIsExist(db, d.DetectionTime)
	if err != nil {
		return err
	}
	if isExist {
		return nil 
	}
	
	result, err := db.Exec(acidRecordSQLString[acidSqliteRecordInsert],
		d.DetectionTime, d.Items[0], d.Items[1], d.Items[2], d.Items[3],
		d.Items[4], d.Items[5])
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

func GetAllAcidRecords(db *sql.DB) ([]*AcidRecord, error) {
	var (
		detectionTime 	time.Time
		items         	[6]float64

		result 		  	[]*AcidRecord
	)

	rows, err := db.Query(acidRecordSQLString[acidSqliteRecordGetAll])
	if err != nil {
		zlog.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&detectionTime, &items[0], &items[1],
			&items[2], &items[3], &items[4], &items[5]); err != nil {
			zlog.Error(err)
			return nil, err
		}
		data := &AcidRecord{
			DetectionTime: detectionTime,
			Items:         items,
		}
		result = append(result, data)
	}

	return result, nil
}

func AcidRecordIsExist(db *sql.DB, detectionTime time.Time) (bool, error) {
	row := db.QueryRow(acidRecordSQLString[acidSqliteRecordIsExist], detectionTime)
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
