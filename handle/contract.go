package handle
import (
	"shop/utils"
	"net/http"
	"errors"
	"shop/config"
	"fmt"
)

func AccountCreatePerson(r *http.Request, param map[string]interface{}) (interface{}, error) {

	data, err :=utils.HttpPostData(config.AddPersonUrl,param)

	if err != nil {
		return nil, err
	}

	if data["errCode"].(float64) != 0 {
		return nil,  errors.New(data["msg"].(string))
	}

	temp := data["data"].(map[string]interface {})

	return temp, nil
}

func CreateOrganizeCommon(r *http.Request, param map[string]interface{}) (interface{}, error) {

	data, err :=utils.HttpPostData(config.AddOrganizeUrl,param)
	fmt.Println(data)
	if err != nil {
		return nil, err
	}

	if data["errCode"].(float64) != 0 {
		return nil,  errors.New(data["msg"].(string))
	}

	temp := data["data"].(map[string]interface {})

	return temp, nil
}


func SilentSign(r *http.Request, param map[string]interface{}) (interface{}, error) {

	data, err :=utils.HttpPostData(config.SilentSignUrl,param)

	if err != nil {
		return nil, err
	}

	if data["errCode"].(float64) != 0 {
		return nil,  errors.New(data["msg"].(string))
	}

	temp := data["data"]

	return temp, nil
}

func UploadTemplateFile(r *http.Request, param map[string]interface{}) (interface{}, error){

	data, err :=utils.GetUploadurl(config.SilentSignUrl,param)

	if err != nil {
		return nil, err
	}

	if data["errCode"].(float64) != 0 {
		return nil,  errors.New(data["msg"].(string))
	}

	temp := data["data"]

	return temp, nil

}









