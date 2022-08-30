package model

import (
	"database/sql"
	"fmt"

	"github.com/silverswords/dielectric/server/util"
	"github.com/silverswords/dielectric/server/zlog"
)

const DielectronTableName = "Dielectron"

const (
	dielectronSqliteRecordCreateTable = iota
	dielectricSqliteCreateTimeIdxIndex
	dielectronSqliteRecordInsert
	dielectronSqliteRecordGetAll
	dielectronSqliteRecordIsExist
)

var (
	dielectronRecordSQLString = []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			detection_time TEXT NOT NULL,
			idx INTEGER NOT NULL DEFAULT 0,
			average REAL NOT NULL DEFAULT 0,
			item1 REAL NOT NULL DEFAULT 0,
			item2 REAL NOT NULL DEFAULT 0,
			item3 REAL NOT NULL DEFAULT 0,
			item4 REAL NOT NULL DEFAULT 0,
			item5 REAL NOT NULL DEFAULT 0,
			item6 REAL NOT NULL DEFAULT 0,
			item7 REAL NOT NULL DEFAULT 0,
			item8 REAL NOT NULL DEFAULT 0,
			item9 REAL NOT NULL DEFAULT 0,
			item10 REAL NOT NULL DEFAULT 0,
			create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`, DielectronTableName),
		fmt.Sprintf(`CREATE UNIQUE INDEX IF NOT EXISTS time_idx ON %s(detection_time, idx)`, DielectronTableName),
		fmt.Sprintf(`INSERT INTO %s(detection_time, idx, average, item1, item2, item3, item4, item5, item6, item7, item8, item9, item10) 
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`, DielectronTableName),
		fmt.Sprintf(`SELECT detection_time, idx, average, item1, item2, item3, item4, item5, item6, item7, item8, item9, item10 FROM %s ORDER BY detection_time DESC, idx ASC`, DielectronTableName),
		fmt.Sprintf(`SELECT COUNT(1) FROM %s WHERE detection_time = ? AND idx = ?`, DielectronTableName),
	}
)

type DielectronRecord struct {
	DetectionTime string   `json:"detection_time,omitempty"`
	Average       float64     `json:"average,omitempty"`
	Index         int         `json:"index,omitempty"`
	Items         [10]float64 `json:"items,omitempty"`
}

func (d *DielectronRecord) String() string {
	return fmt.Sprintf("DateTime:%s\nAvg:%.02f\nNo:%d\nVoltage 1 - 10:%v\n", d.DetectionTime, d.Average, d.Index, d.Items)
}

//CreateTable
func CreateDielectronRecordTable(db *sql.DB) error {
	_, err := db.Exec(dielectronRecordSQLString[dielectronSqliteRecordCreateTable])
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(dielectronRecordSQLString[dielectricSqliteCreateTimeIdxIndex])
	if err != nil {
		panic(err)
	}

	return nil
}

func CheckAndInsertDielectronRecord(db *sql.DB, d *DielectronRecord) error {
	isExist, err := DielectronRecordIsExist(db, d.DetectionTime, d.Index)
	if err != nil {
		return err
	}

	if isExist {
		return nil 
	}
	
	result, err := db.Exec(dielectronRecordSQLString[dielectronSqliteRecordInsert],
		d.DetectionTime, d.Index, d.Average, d.Items[0], d.Items[1], d.Items[2], d.Items[3],
		d.Items[4], d.Items[5], d.Items[6], d.Items[7], d.Items[8], d.Items[9])
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

func GetAllDielectronRecords(db *sql.DB) ([]*DielectronRecord, error) {
	var (
		detectionTime string
		idx int
		average       float64
		items         [10]float64

		result []*DielectronRecord
	)

	rows, err := db.Query(dielectronRecordSQLString[dielectronSqliteRecordGetAll])
	if err != nil {
		zlog.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&detectionTime, &idx, &average, &items[0], &items[1],
			&items[2], &items[3], &items[4], &items[5], &items[6],
			&items[7], &items[8], &items[9]); err != nil {
			zlog.Error(err)
			return nil, err
		}
		data := &DielectronRecord{
			DetectionTime: detectionTime,
			Average:       util.FixPrecision(average),
			Index: idx,
			Items:         items,
		}
		result = append(result, data)
	}

	return result, nil
}

func DielectronRecordIsExist(db *sql.DB, detectionTime string, idx int) (bool, error) {
	row := db.QueryRow(dielectronRecordSQLString[dielectronSqliteRecordIsExist], detectionTime, idx)
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
