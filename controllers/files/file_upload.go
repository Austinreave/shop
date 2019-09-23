package files

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"shop/utils"
	"io/ioutil"
	"log"
	"shop/config"
)

func UploadFile(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	files := []string{"file1"}


	//限制文件到到大小

	//func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser 用来限制接收到的请求的Body的大小的。不同于io.LimitReader，本函数返回一个ReadCloser，返回值的Read方法在读取的数据超过大小限制时会返回非EOF错误，其Close方法会关闭下层的io.ReadCloser接口r。
	r.Body = http.MaxBytesReader(w, r.Body, config.MAX_UPLOAD_SIZE)
	//func (r *Request) ParseMultipartForm(maxMemory int64) error 将请求的主体作为multipart/form-data解析。请求的整个主体都会被解析，得到的文件记录最多maxMemery字节保存在内存，其余部分保存在硬盘的temp文件里。如果必要，ParseMultipartForm会自行调用ParseForm。重复调用本方法是无意义的。
	if err := r.ParseMultipartForm(config.MAX_UPLOAD_SIZE); err != nil {
		log.Printf("file size error: %v", err)
		utils.CheckError(w, err)
		return 
	}


	len := len(files);

	for i := 0; i < len; i++ {
		
		//func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error) FormFile返回以key为键查询r.MultipartForm字段得到结果中的第一个文件和它的信息。如果必要，本函数会隐式调用ParseMultipartForm和ParseForm。
		file, imgHead, err := r.FormFile(files[i])

		if err != nil {
			log.Printf("Error when try to get file: %v", err)
			utils.CheckError(w, err)
			return 
		}

		data, err := ioutil.ReadAll(file) //func ReadAll(r io.Reader) ([]byte, error)

		if err != nil {
			log.Printf("Read file error: %v", err)
			utils.CheckError(w, err)
		}

		//验证文件类型
		filetype := http.DetectContentType(data) //func DetectContentType(data []byte) string 用于确定数据的Content-Type

		err = utils.VerifyFileType(filetype)

		if err != nil {
			log.Printf("file type error: %v", err)
			utils.CheckError(w, err)
			return 
		}
		//func WriteFile(filename string, data []byte, perm os.FileMode) error 函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件
		err = ioutil.WriteFile(config.FILE_DIR+imgHead.Filename, data, 0666)//写入文件
		if err != nil {
			log.Printf("Write file error: %v", err)
			utils.CheckError(w, err)
			return
		}
	}

	utils.Success(w, len)
	
}




















