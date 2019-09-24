package contract

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"shop/handle"
	"shop/utils"
)

func AccountCreatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"name","idNo","idType"}

	param, err := utils.AcceptData(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.AccountCreatePerson(r,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}

func CreateOrganizeCommon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"creatorId","name","organCode","organType","legalName","legalIdNo"}

	param, err := utils.AcceptData(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.CreateOrganizeCommon(r,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}


func SilentSign(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"grantAccountId"}

	param, err := utils.AcceptData(r, s...)
	
	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.SilentSign(r,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}

func UploadTemplateFile(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	s := []string{"filePath","fileName"}
	//接受用户输入
	param, err := utils.AcceptData(r, s...)
	
	if err != nil {
		utils.CheckError(w, err)
		return
	}
	//接收e签宝
	d,err := handle.UploadTemplateFile(r,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}










