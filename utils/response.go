package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

type JsonResult struct {
	Code    int
	Message string
	Data    interface{}
}

func Success(w http.ResponseWriter, data interface{}) {

	var result JsonResult
	result.Message = "SUCCESS"
	result.Code = 0
	result.Data = data
	bytes, _ := json.Marshal(result)

	io.WriteString(w, string(bytes))//响应页面
}

func CheckError(w http.ResponseWriter, err error) {
	var result JsonResult
	result.Message = fmt.Sprintf("%s", err)
	result.Code = 1
	result.Data = nil
	bytes, _ := json.Marshal(result)
	io.WriteString(w, string(bytes))//响应页面
}




