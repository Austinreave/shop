package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
 	"io/ioutil"
	"encoding/json"
	"shop/config"
	"os"
	"errors"
	"shop/languages"
	"fmt"
)

func QueryArrays(db *sql.DB, q string, args ...interface{}) ([]map[string]interface{}, error) {


	//func (db *DB) Query(query string, args ...interface{}) (*Rows, error) Query执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数
	rows, err := db.Query(q, args...)

	if err != nil {
		return nil, err
	}

	// Get column names
	columns, err := rows.Columns()//func (rs *Rows) Columns() ([]string, error) Columns返回列名。如果Rows已经关闭会返回错误

	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))//RawBytes是一个字节切片，保管对内存的引用，为数据库自身所使用。在Scaner接口的Scan方法写入RawBytes数据后，该切片只在限次调用Next、Scan或Close方法之前合法

	scanArgs := make([]interface{}, len(values))
	
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// 最终要返回的data
	data := make([]map[string]interface{}, 0)
	// Fetch rows
	for rows.Next() {//Next准备用于Scan方法的下一行结果
		err = rows.Scan(scanArgs...)//func (rs *Rows) Scan(dest ...interface{}) error Scan将当前行各列结果填充进dest指定的各个值中
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

	if err != nil {
		return nil, err
	}
	
    return temp,err

}

func SendHttpPUT(url, contentBase64Md5  string, fileContent []byte) (map[string]interface{}, error) {

    req, _ := http.NewRequest("PUT", url, strings.NewReader(string(fileContent)))

    req.Header.Add("Content-Type","application/pdf")
    req.Header.Add("Content-Md5", contentBase64Md5)

    resp, err := (&http.Client{}).Do(req)

	if err != nil {
		return nil, err
	}

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
	
    temp := make(map[string]interface{}, 0)

	err = json.Unmarshal(body, &temp)

	if err != nil {
		return nil, err
	}
	
    return temp,err
}



func GetUploadurl(url string, param map[string]interface{}) (map[string]interface{}, error) {

	filePath, ok := param["filePath"].(string)

 	if ok != true {
	   return nil, errors.New(languages.AssertionFailure)
	}


	f, err := os.Open(filePath)
	if err != nil {
	   return nil, err
	}

 	defer f.Close()

	fileContent, err:= ioutil.ReadAll(f)

	if err != nil {
	   return nil, err
	}

	fileSize := len(fileContent)

    fileMd5, err := GetContentBase64Md5(filePath)

	if err != nil {
	   return nil, err
	}

 	param["fileSize"] = fileSize
 	param["contentType"] = config.ContentType
 	param["contentMd5"] = fileMd5

 	temp, err :=HttpPostData(url,param)

 	fmt.Println(temp)

 	if err != nil {
	   return nil, err
	}

	contentMd5, ok := fileMd5.(string)


 	if ok != true {
	   return nil, errors.New(languages.AssertionFailure)
	}

	data, ok := temp["data"].(map[string]interface {})

 	if ok != true {
	   return nil, errors.New(languages.AssertionFailure)
	}

	uploadUrl, ok := data["uploadUrl"].(string)

 	if ok != true {
	   return nil, errors.New(languages.AssertionFailure)
	}

	result, err := SendHttpPUT(uploadUrl, contentMd5, fileContent)

	if err != nil {
	   return nil, err
	}

	return result,nil

}




