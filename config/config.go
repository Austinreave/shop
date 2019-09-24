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

//e签宝
const(
    PROJECT_ID = "4438758767"
    PROJECT_SECRET = "cc7fb6e1278990edf87f3e4f0f9c3a51"
    API_HOST = "https://smlo.tsign.cn/opentreaty-service/"
    AddPersonUrl = API_HOST+"account/create/person" //创建个人账户地址
    AddOrganizeUrl = API_HOST+"account/create/organize/common";//创建企业账户地址
    SilentSignUrl = API_HOST+"account/platform/silentsign";//设置静默签署授权地址
    GetUploadurl = API_HOST+"file/uploadurl";//获取文件直传地址
)

//文件类型
const(
    ContentType = "application/pdf"
)
