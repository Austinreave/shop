package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
	"fmt"
 	"io/ioutil"
	"encoding/json"
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


func HttpPostData(args ...interface{}){

    byte, _ := json.Marshal(args)

    req, _ := http.NewRequest("POST", "http://abced.com/" + "/user/false/lsj", strings.NewReader(string(byte)))
    req.Header.Set("token", "00998ecf8427e")
    resp, err := (&http.Client{}).Do(req)
    if err != nil {
  		 fmt.Println(err)

    }
    defer resp.Body.Close()
    respByte, _ := ioutil.ReadAll(resp.Body)

    fmt.Println(respByte)

	//模拟Post请求
	// $return_content = $this->http_post_data($this->addPersonUrl,$data,$this->projectID,$this->projectSecret);
	// function http_post_data($url, $data, $projectid, $projectSecret) {
	//     $ch = curl_init();
	//     curl_setopt($ch, CURLOPT_POST, 1);
	//     curl_setopt($ch, CURLOPT_URL, $url);
	//     curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
	//     curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);  // 跳过检查
	//     curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);  // 跳过检查
	//     curl_setopt($ch, CURLOPT_HTTPHEADER, array("X-Tsign-Open-App-Id:".$projectid, "X-Tsign-Open-App-Secret:".$projectSecret, "Content-Type:application/json" ));
	//     ob_start();
	//     curl_exec($ch);
	//     $return_content = ob_get_contents();
	//     ob_end_clean();
	//     $return_code = curl_getinfo($ch, CURLINFO_HTTP_CODE);
	//     return array($return_code, $return_content);
	// }
}

