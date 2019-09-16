package contract

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"shop/handle"
	"shop/utils"
	"errors"
)

func AccountCreatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"name","idNo","idType"}

	param, err := utils.FindParam(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	handle.AccountCreatePerson(param)

	// if err != nil{
	// 	utils.CheckError(w, errors.New("获取失败"))
	// 	return
	// }

	// utils.Success(w, d)

}