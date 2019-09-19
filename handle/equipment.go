package handle

import (
	"shop/utils"
	"net/http"
	"shop/database"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func GetEquipmentList(r *http.Request, param []interface{}) ([]map[string]interface{}, error) {

	sql := "SELECT * FROM ecm_equipment WHERE user_name LIKE CONCAT('%',?,'%') OR equipment_name LIKE CONCAT('%',?,'%') LIMIT ?,?"

	d, err := utils.QueryArrays(database.DbConn, sql, param...)

	if err != nil {
		return nil, err
	}
	
	return d, nil
}

func GetEquipment(r *http.Request, p httprouter.Params,equipment_id interface{}) ([]map[string]interface{}, error) {

	sql := fmt.Sprintf("SELECT * FROM ecm_equipment WHERE equipment_id = ?")

	d, err := utils.QueryArrays(database.DbConn, sql, equipment_id)

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

func DeleteEquipment(r *http.Request, p httprouter.Params, equipment_id interface{}) (int64, error) {

	sql := fmt.Sprintf("DELETE FROM ecm_equipment WHERE equipment_id = ?")

	result, err := utils.ExecuteSQL(database.DbConn, sql, equipment_id)

	if err != nil {

		return 0, err
	}

	return result.RowsAffected()
}





