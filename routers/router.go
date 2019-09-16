package routers

import(
    "github.com/julienschmidt/httprouter"
	"shop/controllers/equipment"
	"shop/controllers/files"
	"shop/config"
	"net/http"
)

func Loading(router *httprouter.Router) *httprouter.Router {
	
	//设备
	router.GET("/equipment",equipment.GetList)
	router.GET("/equipment/:id",equipment.GetDetail)
	router.POST("/equipment",equipment.Post)
	router.PUT("/equipment",equipment.Put)
	router.DELETE("/equipment/:id",equipment.Delete)

	//文件
	router.POST("/file",files.UploadFile)
	router.ServeFiles("/static/*filepath",http.Dir(config.FILE_DIR))

	//e签宝
	router.POST("/contract/account_create_person",contract.AccountCreatePerson)//创建个人账号



	return router
}