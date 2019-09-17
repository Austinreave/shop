package contract

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"shop/handle/contract"
	"shop/utils"
)

func AccountCreatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"name","idNo","idType"}

	param, err := utils.AcceptData(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := contract.AccountCreatePerson(r,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w, d)

}