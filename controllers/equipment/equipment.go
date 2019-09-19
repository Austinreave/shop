package equipment

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"shop/handle"
	"shop/utils"
)

func GetList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"user_name","equipment_name","offset","psize"}

	param, err := utils.FindParam(r,s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.GetEquipmentList(r, param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}
	utils.Success(w, d)
}

func GetDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	equipment_id, err := utils.FindOneParam(r,"equipment_id")


	if err != nil {
		utils.CheckError(w, err)
	}

	d,err := handle.GetEquipment(r, p, equipment_id)

	if err != nil{
		utils.CheckError(w, err)
		return
	}
	utils.Success(w, d)
}


func Post(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"area_code","create_time","equipment_brand","equipment_details","equipment_name","equipment_status","update_time","user_id","user_name"}

	param, err := utils.FindParam(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.AddEquipment(r, p, param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w, d)

}


func Put(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	s := []string{"area_code","create_time","equipment_brand","equipment_details","equipment_name","equipment_status","update_time","user_id","user_name","equipment_id"}

	param, err := utils.FindParam(r, s...)

	if err != nil {
		utils.CheckError(w, err)
	}

	d,err := handle.UpdateEquipment(r, p, param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w, d)
}


func Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	equipment_id, err := utils.FindOneParam(r,"equipment_id")

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.DeleteEquipment(r, p, equipment_id)

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	utils.Success(w, d)
	
}




