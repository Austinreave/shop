package utils

import(
 	"net/http"
 	"shop/config"
 	"errors"
 	"shop/languages"
)

func ValidateAuth(r *http.Request) error {

	md5Str := md5V(config.SecretKey)

	token := r.Header.Get("token")

	if md5Str != token {
		return errors.New(languages.ValidationFails)
	}
	
	return nil

}












