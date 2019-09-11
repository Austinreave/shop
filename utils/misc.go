package utils

import (
 	"crypto/md5"
 	"encoding/hex"
 	"errors"
)

func md5V(str string) string  {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func VerifyFileType(filetype string) error {
	switch filetype {
		case "image/jpeg", "image/jpg":
			return nil 
		case "image/gif", "image/png":
			return nil 
		case "application/pdf":
			return nil 
		default:
			return errors.New("文件类型错误")
	}

}


