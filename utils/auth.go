package utils

import(
 	"net/http"
 	"shop/config"
 	"errors"
 	"shop/languages"
)

func ValidateAuth(r *http.Request) error {

	md5Str := md5V(config.SecretKey)

	token := r.Header.Get("token") //获取header头信息

	if md5Str != token {
		return errors.New(languages.ValidationFails)//使用字符串创建一个错误,请类比fmt包的Errorf方法，差不多可以认为是New(fmt.Sprintf(...))
	}
	
	return nil
}












