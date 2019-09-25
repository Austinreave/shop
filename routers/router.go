package routers

import(
    "github.com/julienschmidt/httprouter"
	"shop/config"
	"net/http"
	"shop/controllers/equipment"
	"shop/controllers/files"
	"shop/controllers/contract"
)

func Loading(router *httprouter.Router) *httprouter.Router {
	
	//设备
	router.GET("/equipment",equipment.GetList)
	router.GET("/equipment/detail",equipment.GetDetail)
	router.POST("/equipment",equipment.Post)
	router.PUT("/equipment",equipment.Put)
	router.DELETE("/equipment",equipment.Delete)

	//文件
	router.POST("/file",files.UploadFile)
	router.ServeFiles("/static/*filepath",http.Dir(config.FILE_DIR))

	//e签宝
	router.POST("/contract/account_create_person",contract.AccountCreatePerson)//创建个人账号
	router.POST("/contract/create_organize_common",contract.CreateOrganizeCommon)//创建企业账号
	router.POST("/contract/silent_sign",contract.SilentSign)//创建企业账号
	router.POST("/contract/upload_template_file",contract.UploadTemplateFile)//上传文件模板
	router.POST("/contract/create_file_key",contract.CreateFileKey)//创建模板
	router.POST("/contract/create_template",contract.CreateTemplate)//创建合同
	router.POST("/contract/create_process",contract.CreateProcess)//创建合同签署流程
	router.POST("/contract/user_sign_task",contract.UserSignTask)//发起用户自动签署

	return router
}