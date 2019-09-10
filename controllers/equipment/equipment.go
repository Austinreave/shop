package equipment

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"shop/handle"
	"shop/utils"
	"errors"
)

func GetList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	offset, psize, err := utils.GetPageParam(r)

	if err != nil {
		utils.CheckError(w, err)
	}

	d,err := handle.GetEquipmentList(r, offset.(string), psize.(string))

	if err != nil{
		utils.CheckError(w, errors.New("获取失败"))
		return
	}
	utils.Success(w, d)
}

func GetDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	id, err := utils.GetOneParam(p)

	if err != nil {
		utils.CheckError(w, err)
	}

	d,err := handle.GetEquipment(r, p, id.(string))

	if err != nil{
		utils.CheckError(w, errors.New("获取失败"))
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
		utils.CheckError(w, errors.New("添加失败"))
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
		utils.CheckError(w, errors.New("修改失败"))
		return
	}

	utils.Success(w, d)
}


func Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, err := utils.GetOneParam(p)

	if err != nil {
		utils.CheckError(w, err)
	}

	d,err := handle.DeleteEquipment(r, p, id.(string))

	if err != nil {
		utils.CheckError(w, errors.New("删除失败"))
		return
	}

	utils.Success(w, d)
	
}




