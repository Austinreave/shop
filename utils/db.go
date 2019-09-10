package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func QueryArrays(db *sql.DB, q string, args ...interface{}) ([]map[string]interface{}, error) {

	rows, err := db.Query(q, args...)


	if err != nil {
		return nil, err
	}

	// Get column names
	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// 最终要返回的data
	data := make([]map[string]interface{}, 0)
	// Fetch rows
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		//每行数据
		row := make(map[string]interface{})
		for k, v := range values {
			key := columns[k]
			row[key] = string(v)
		}
		//放入结果集
		data = append(data, row)
	}

	return data, nil
}

func ExecuteSQL(db *sql.DB, q string, args ...interface{}) (sql.Result, error) {

	stmt, err := db.Prepare(q)

	if err != nil {
		return nil, err
	}
	
	res, err := stmt.Exec(args...)

	if err != nil {
		return nil, err
	}

	return res, err
}

