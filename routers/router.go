package routers

import(
    "github.com/julienschmidt/httprouter"
	"shop/controllers/equipment"
)

func Loading(router *httprouter.Router) *httprouter.Router {
	//设备
	router.GET("/equipment",equipment.GetList)
	router.GET("/equipment/:id",equipment.GetDetail)
	router.POST("/equipment",equipment.Post)
	router.PUT("/equipment",equipment.Put)
	router.DELETE("/equipment/:id",equipment.Delete)

	return router
}