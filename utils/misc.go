package utils

import (
 	"crypto/md5" //md5包实现了MD5哈希算法
 	"encoding/hex" //hex包实现了16进制字符表示的编解码
 	"errors"
 	"shop/languages"
 	"io"
 	"encoding/base64"
 	"os"
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


func GetContentBase64Md5(filePath string) (interface{}, error) {

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := md5.New()

	if _, err = io.Copy(hash, file); err != nil {
		return nil, err
	}

	hashInBytes := hash.Sum(nil)[:16]

    //计算文件的Content-MD5
    contentBase64Md5 := base64.StdEncoding.EncodeToString(hashInBytes)
    
    return contentBase64Md5,nil
}


