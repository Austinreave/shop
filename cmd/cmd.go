package cmd

import(
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "shop/routers"
    "shop/utils"
)

//用来做中间件使用
type middleWareHandler struct {
	r *httprouter.Router //定义一个httprouter.Router用来继承
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r //赋值
	return m
}

//实现ServeHTTP接口（使用 duck type模式）
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := utils.ValidateAuth(r)
	if err != nil {
		utils.CheckError(w, err)
		return 
	}
	m.r.ServeHTTP(w, r)//执行具体方法
}

func ShopStart() error {
	mh := NewMiddleWareHandler(routers.Loading(httprouter.New()))
	err := http.ListenAndServe(":80",mh)
	fmt.Println(err.Error())
	return nil
}






