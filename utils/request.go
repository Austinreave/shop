package utils
import (
 	"net/http"
 	"errors"
 	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func FindParam(r *http.Request, s ...string) ([]interface{}, error) {

	body, _ := ioutil.ReadAll(r.Body)

	h := make(map[string]interface{}, 0)

	err := json.Unmarshal(body, &h)
	if err != nil {
		return nil, err
	}
	
	if len(s) < 0 {
		return nil, errors.New("slice len is 0")
	}

	param := s[0]

	if _, ok := h[param]; !ok {
		return nil, errors.New("slices find not param")
	}

	var temp []interface{}

	for _, v := range s {
		temp = append(temp, h[v])
	}

	return temp, nil
}

func GetOneParam(p httprouter.Params) (interface{}, error) {

	param := p.ByName("id")

	if param == "" {
		return nil, errors.New("parameter error")
	}

	return param, nil
}


func GetPageParam(r *http.Request) (interface{}, interface{} ,error) {

	offset :=r.Form["offset"][0]

	psize :=r.Form["psize"][0]

	if offset == "" {
		return nil, nil,errors.New("parameter error")
	}
	if psize == "" {
		return nil, nil,errors.New("parameter error")
	}

	return offset, psize, nil
}


func GetSearchParam(r *http.Request,s ...string ) (interface{},error) {
	
	body, _ := ioutil.ReadAll(r.Body)
	h := make(map[string]interface{}, 0)
	err := json.Unmarshal(body, &h)

	fmt.Println(h)
	fmt.Println(err)
	return nil, nil


}









