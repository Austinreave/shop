package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
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


func GetUploadurl(filePath,fileName,url string){

    // $fileSize = strlen(file_get_contents(filePath));
    // $contentType = "application/pdf";
    // $contentMd5 = $this->getContentBase64Md5(filePath);
    // $arr = array('fileName'=>$fileName,'fileSize'=>$fileSize,'contentType'=>$contentType,'contentMd5'=>$contentMd5);
    // //将数组转成json字符串（JSON_UNESCAPED_SLASHES 此参数是为了不让application/pdf 中的“/”被转义掉）
    // $data = json_encode($arr,JSON_UNESCAPED_SLASHES);
    // $result = $this->doPost(Url,$data,config.PROJECT_ID,config.PROJECT_SECRET);
    // return $result;
}



func HttpPostData(url string,data map[string]interface{}) (map[string]interface{}, error){

    byte, _ := json.Marshal(data)

    req, _ := http.NewRequest("POST", url, strings.NewReader(string(byte)))

    req.Header.Add("X-Tsign-Open-App-Id", config.PROJECT_ID)
    req.Header.Add("X-Tsign-Open-App-Secret", config.PROJECT_SECRET)
    req.Header.Add("Content-Type","application/json")

    resp, err := (&http.Client{}).Do(req)	

	if err != nil {
		return nil, err
	}

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    temp := make(map[string]interface{}, 0)

	err = json.Unmarshal(body, &temp)
	
    return temp,err

}

