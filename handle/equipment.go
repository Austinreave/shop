package handle

import (
	"shop/utils"
	"net/http"
	"shop/database"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func GetEquipmentList(r *http.Request, offset string, psize string) ([]map[string]interface{}, error) {

	sql := fmt.Sprintf("SELECT * FROM ecm_equipment LIMIT "+offset+","+psize+"")
	
	d, err := utils.QueryArrays(database.DbConn, sql)

	if err != nil {
		return nil, err
	}
	
	return d, nil
}

func GetEquipment(r *http.Request, p httprouter.Params,id string) ([]map[string]interface{}, error) {

	sql := fmt.Sprintf("SELECT * FROM ecm_equipment WHERE equipment_id = "+id+"")

	d, err := utils.QueryArrays(database.DbConn, sql)

	if err != nil {
		return nil, err
	}
	
	return d, nil
}

func AddEquipment(r *http.Request, p httprouter.Params, param []interface{}) (int64, error) {

	sql := fmt.Sprintf("INSERT ecm_equipment SET area_code = ?,create_time = ?,equipment_brand = ?,equipment_details = ?, equipment_name = ?, equipment_status =?, update_time = ?, user_id = ?,user_name = ?")
	
	res, err := utils.ExecuteSQL(database.DbConn, sql, param...)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateEquipment(r *http.Request, p httprouter.Params, param []interface{}) (int64, error) {

	sql := fmt.Sprintf("UPDATE ecm_equipment SET area_code = ?,create_time = ?,equipment_brand = ?,equipment_details = ?, equipment_name = ?, equipment_status =?, update_time = ?, user_id = ?,user_name = ? WHERE equipment_id = ?")
	
	result, err := utils.ExecuteSQL(database.DbConn, sql, param...)

	if err != nil {

		return 0, err
	}

	return result.RowsAffected()
}

func DeleteEquipment(r *http.Request, p httprouter.Params, id string) (int64, error) {

	sql := fmt.Sprintf("DELETE FROM ecm_equipment WHERE equipment_id = ?")

	result, err := utils.ExecuteSQL(database.DbConn, sql, id)

	if err != nil {

		return 0, err
	}

	return result.RowsAffected()
}





