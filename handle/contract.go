package handle
import (
	"shop/utils"
	"net/http"
	"errors"
)

func AccountCreatePerson(r *http.Request, param map[string]interface{}) (interface{}, error) {

	data, err :=utils.HttpPostData(param)

	if err != nil {
		return nil, err
	}

	if data["errCode"].(float64) != 0 {
		return nil,  errors.New(data["msg"].(string))
	}

	temp := data["data"].(map[string]interface {})

	return temp, nil
}
