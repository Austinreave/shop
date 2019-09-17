package contract
import (
	"shop/utils"
	"net/http"
	"errors"
	"fmt"
)

func AccountCreatePerson(r *http.Request, param map[string]interface{}) (map[string]interface{}, error) {

	data, err :=utils.HttpPostData(param)

	if err != nil {
		return nil, err
	}

	errCode := data["errCode"]
	if errCode != 0 {
		return nil, errors.New(data["msg"].(string))
	}

	// temp := make(map[string]interface{}, 0)

	fmt.Println(data["data"])

	return data, nil
}
