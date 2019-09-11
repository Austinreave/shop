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

	r.Body = http.MaxBytesReader(w, r.Body, config.MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(config.MAX_UPLOAD_SIZE); err != nil {
		log.Printf("file size error: %v", err)
		utils.CheckError(w, err)
		return 
	}

	len := len(files);

	for i := 0; i < len; i++ {
		
		//获取文件
		file, imgHead, err := r.FormFile(files[i])

		if err != nil {
			log.Printf("Error when try to get file: %v", err)
			utils.CheckError(w, err)
			return 
		}

		data, err := ioutil.ReadAll(file)

		if err != nil {
			log.Printf("Read file error: %v", err)
			utils.CheckError(w, err)
		}

		//验证文件类型
		filetype := http.DetectContentType(data)

		err = utils.VerifyFileType(filetype)

		if err != nil {
			log.Printf("file type error: %v", err)
			utils.CheckError(w, err)
			return 
		}

		err = ioutil.WriteFile(config.FILE_DIR+imgHead.Filename, data, 0666)
		if err != nil {
			log.Printf("Write file error: %v", err)
			utils.CheckError(w, err)
			return
		}
	}

	utils.Success(w, len)
	
}




















