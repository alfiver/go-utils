package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"sync"
)

var (
	db    *sql.DB
	dbMtx sync.Mutex
)

func Init(f string) error {
	if err := os.MkdirAll(path.Dir(f), 0755); err != nil {
		return errors.New("sqlite dir mkdir failed! " + err.Error())
	}
	dbR, err := sql.Open("sqlite3", f)
	if err != nil {
		return errors.New("sqlite open failed! " + err.Error())
	}
	db = dbR
	return nil
}
func Insert(sql string, vals ...any) (int64, error) {
	dbMtx.Lock()
	defer dbMtx.Unlock()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, errors.New(sql + " prepare failed! " + err.Error())
	}
	defer stmt.Close()
	res, err := stmt.Exec(vals...)
	if err != nil {
		return 0, errors.New(sql + " exec failed! " + err.Error())
	}
	return res.LastInsertId()
}
func RawExec(sql string, vals ...any) error {
	dbMtx.Lock()
	defer dbMtx.Unlock()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(vals...)
	if err != nil {
		return err
	}
	return nil
}
func RawQuery(query string, vals ...any) ([]map[string]string, error) {
	dbMtx.Lock()
	defer dbMtx.Unlock()
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(vals...)
	if err != nil {
		return nil, err
	}
	cols, _ := rows.Columns()
	results := make([]map[string]string, 0)
	for rows.Next() {
		values := make([]sql.RawBytes, len(cols))
		scans := make([]any, len(cols))
		for i := range values {
			scans[i] = &values[i]
		}
		if err = rows.Scan(scans...); err == nil {
			rm := make(map[string]string)
			for j := range values {
				rm[cols[j]] = string(values[j])
			}
			results = append(results, rm)
		}
	}
	return results, nil
}
func RawQueryOne(query string, vals ...any) (map[string]string, error) {
	results, err := RawQuery(query, vals...)
	if err != nil {
		return nil, err
	}
	if len(results) > 0 {
		return results[0], nil
	}
	return nil, nil
}
