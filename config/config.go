package config

//签名秘钥
const SecretKey string = "shop123"

//mysql 
const(
	Conn = "root:abc123456@tcp(39.96.85.45:3307)/dbc_shop?charset=utf8"
)

//文件
const (
	FILE_DIR = "./static/"
	MAX_UPLOAD_SIZE = 1024 * 1024 * 8 //50MB
)