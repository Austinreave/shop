package utils
import (
 	"net/http"
 	"errors"
 	"io/ioutil"
	"encoding/json"
	"shop/languages"
)

func AcceptData(r *http.Request, s ...string) (map[string]interface{}, error) {

	body, _ := ioutil.ReadAll(r.Body)

	temp := make(map[string]interface{}, 0)

	data := make(map[string]interface{}, 0)

	err := json.Unmarshal(body, &temp)

	if err != nil {
		return nil, err
	}

	//提取需要的数据
	for _, v := range s {
		if temp[v] == nil {
			return nil, errors.New(languages.ParameterAcquisitionFailed)
		}
		data[v] = temp[v]
	}
	return data, nil
}


func FindParam(r *http.Request, s ...string) ([]interface{}, error) {

	body, _ := ioutil.ReadAll(r.Body)

	h := make(map[string]interface{}, 0)

	err := json.Unmarshal(body, &h)

	if err != nil {
		return nil, err
	}
	
	if len(s) < 0 {
		return nil, errors.New(languages.SliceError)
	}

	param := s[0]

	if _, ok := h[param]; !ok {
		return nil, errors.New(languages.SliceFindError)
	}

	var temp []interface{}

	for _, v := range s {
		temp = append(temp, h[v])
	}

	return temp, nil
}

func FindOneParam(r *http.Request, s string) (interface{}, error){

	body, _ := ioutil.ReadAll(r.Body)

	h := make(map[string]interface{}, 0)

	err := json.Unmarshal(body, &h)

	if err != nil {
		return nil, err
	}

	if len(s) < 0 {
		return nil, errors.New(languages.SliceError)
	}
	
	return h[s], nil
}










