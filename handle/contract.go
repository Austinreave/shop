package handle
import (
	"shop/utils"
	"errors"
	"shop/languages"
)

func CreateData(url string, param map[string]interface{}) (interface{}, error) {

	data, err :=utils.HttpPostData(url,param)

	if err != nil {
		return nil, err
	}

	errCode, ok := data["errCode"].(float64)

	if ok != true {
		return nil,  errors.New(languages.AssertionFailure)
	}

	if errCode != 0 {
		return nil,  errors.New(data["msg"].(string))
	}

	temp, ok := data["data"].(map[string]interface {})

	if ok != true {
		return nil,  errors.New(languages.AssertionFailure)
	}

	return temp, nil
}









