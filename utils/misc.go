package utils

import (
 	"crypto/md5" //md5包实现了MD5哈希算法
 	"encoding/hex" //hex包实现了16进制字符表示的编解码
 	"errors"
 	"shop/languages"
 	"fmt"
)

func md5V(str string) string  {
	
    h := md5.New()//func New() hash.Hash
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))//将数据src编码为字符串s。
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
			return errors.New(languages.FileTypeError)//使用字符串创建一个错误,请类比fmt包的Errorf方法，差不多可以认为是New(fmt.Sprintf(...))
	}

}


