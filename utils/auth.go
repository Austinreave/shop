package utils

import(
 	"net/http"
 	"shop/config"
 	"errors"
)

func ValidateAuth(r *http.Request) error {
	md5Str := md5V(config.SecretKey)
	token := r.Header.Get("token")
	if md5Str == token {
		return errors.New("验证失败")
	}
	r.ParseForm()
	return nil
}












