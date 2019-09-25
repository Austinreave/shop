package contract

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"shop/handle"
	"shop/utils"
	"shop/config"
)

func AccountCreatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"name","idNo","idType"}

	param, err := utils.AcceptData(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.CreateData(config.AddPersonUrl,param)

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

	d,err := handle.CreateData(config.AddOrganizeUrl,param)

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

	d,err := handle.CreateData(config.SilentSignUrl,param)

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
	d,err := handle.CreateData(config.GetUploadurl,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}

func CreateFileKey(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	s := []string{"fileKey","templateName"}
	//接受用户输入
	param, err := utils.AcceptData(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}
	//接收e签宝
	d,err := handle.CreateData(config.CreateByFileKeyUrl,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}

func CreateTemplate (w http.ResponseWriter, r *http.Request, p httprouter.Params){

	s := []string{"name","templateId","simpleFormFields"}
	//接受用户输入
	param, err := utils.AcceptData(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}
	//接收e签宝
	d,err := handle.CreateData(config.CreateByTemplateUrl,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}

func CreateProcess (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	s := []string{"businessScene","docId"}
	//接受用户输入
	param, err := utils.AcceptData(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}
	//接收e签宝
	d,err := handle.CreateData(config.AddProcessUrl,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}

func UserSignTask (w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"flowId","accountId","posList"}

	//接受用户输入
	param, err := utils.AcceptData(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}
	//接收e签宝
	d,err := handle.CreateData(config.UserSignTaskUrl,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}












