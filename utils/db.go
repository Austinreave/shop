package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
	"fmt"
 	"io/ioutil"
	"encoding/json"
	"shop/config"
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


func HttpPostData(data interface{}){

    byte, _ := json.Marshal(data)
    req, _ := http.NewRequest("POST", config.AddPersonUrl, strings.NewReader(string(byte)))

    req.Header.Add("X-Tsign-Open-App-Id", config.PROJECT_ID)
    req.Header.Add("X-Tsign-Open-App-Secret", config.PROJECT_SECRET)
    req.Header.Add("Content-Type","application/json")

    resp, err := (&http.Client{}).Do(req)	

    if err != nil {
  		 fmt.Println(err)
    }

    defer resp.Body.Close()

    respByte, _ := ioutil.ReadAll(resp.Body)

    fmt.Println(string(respByte))

}

