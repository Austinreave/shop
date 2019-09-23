package contract

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"shop/handle"
	"shop/utils"
)

func AccountCreatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"name","idNo","idType"}

	param, err := utils.AcceptData(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.AccountCreatePerson(r,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}

func CreateOrganizeCommon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"creatorId","name","organCode","organType","legalName","legalIdNo"}

	param, err := utils.AcceptData(r, s...)

	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.CreateOrganizeCommon(r,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}


func SilentSign(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := []string{"grantAccountId"}

	param, err := utils.AcceptData(r, s...)
	
	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.SilentSign(r,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}

func UploadTemplateFile(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	s := []string{"filePath","fileName"}

	param, err := utils.AcceptData(r, s...)
	
	if err != nil {
		utils.CheckError(w, err)
		return
	}

	d,err := handle.UploadTemplateFile(r,param)

	if err != nil{
		utils.CheckError(w, err)
		return
	}

	utils.Success(w,d)

}


// //获取文件直传地址、上传文件
// function file_upload($filePath,$fileName)
// {	
// 	//fileKey  uploadUrl	
// 	$dataRes = $this->getUploadurl($filePath,$this->get_Uploadurl,$this->projectID,$this->projectSecret,$fileName);

// 	$uploadRes=$this->uploadFile($filePath,$dataRes['data']['uploadUrl']);


// 	return $dataRes;
// }

// //获取文件直传地址
// function getUploadurl($filePath,$get_Uploadurl,$projectID,$projectSecret,$fileName){

//     $fileSize = strlen(file_get_contents($filePath));
//     $contentType = "application/pdf";
//     $contentMd5 = $this->getContentBase64Md5($filePath);
//     $arr = array('fileName'=>$fileName,'fileSize'=>$fileSize,'contentType'=>$contentType,'contentMd5'=>$contentMd5);
//     //将数组转成json字符串（JSON_UNESCAPED_SLASHES 此参数是为了不让application/pdf 中的“/”被转义掉）
//     $data = json_encode($arr,JSON_UNESCAPED_SLASHES);
//     $result = $this->doPost($get_Uploadurl,$data,$projectID,$projectSecret);
//     return $result;
// }

// //上传文件
// function uploadFile($filePath,$uploadUrl){
//     $fileContent = file_get_contents($filePath);
//     $contentMd5 = $this->getContentBase64Md5($filePath);
//     $status = $this->sendHttpPUT($uploadUrl, $contentMd5, $fileContent);
//     if ($status == 200) {
//         return ['status' => 'success', 'code' => $status, 'msg' => '文件上传成功'];
//     }
//     return ['status' => 'error', 'code' => $status, 'msg' => '文件上传失败'];
// }

// /**
// 获取文件的Content-MD5
// 原理：1.先计算MD5加密的二进制数组（128位）
// 2.再对这个二进制进行base64编码（而不是对32位字符串编码）
// */
// function getContentBase64Md5($filePath){
//     //获取文件MD5的128位二进制数组
//     $md5file = md5_file($filePath,true);
//     //计算文件的Content-MD5
//     $contentBase64Md5 = base64_encode($md5file);
//     return $contentBase64Md5;
// }

// /**
//  *
//  * 模拟发送PUT方式请求
//  * @param $url
//  * @param $contentBase64Md5
//  * @param $fileContent
//  * @return mixed|string
//  * @author Ayz
//  */
// function sendHttpPUT($url, $contentBase64Md5, $fileContent)
// {
//     $header = [
//         'Content-Type:application/pdf',
//         'Content-Md5:' . $contentBase64Md5
//     ];

//     $status = '';
//     $curl_handle = curl_init();
//     curl_setopt($curl_handle, CURLOPT_URL, $url);
//     curl_setopt($curl_handle, CURLOPT_FILETIME, true);
//     curl_setopt($curl_handle, CURLOPT_FRESH_CONNECT, false);
//     curl_setopt($curl_handle, CURLOPT_HEADER, true); // 输出HTTP头 true
//     curl_setopt($curl_handle, CURLOPT_RETURNTRANSFER, true);
//     curl_setopt($curl_handle, CURLOPT_TIMEOUT, 5184000);
//     curl_setopt($curl_handle, CURLOPT_CONNECTTIMEOUT, 120);
//     curl_setopt($curl_handle, CURLOPT_SSL_VERIFYPEER, false);
//     curl_setopt($curl_handle, CURLOPT_SSL_VERIFYHOST, false);

//     curl_setopt($curl_handle, CURLOPT_HTTPHEADER, $header);
//     curl_setopt($curl_handle, CURLOPT_CUSTOMREQUEST, 'PUT');

//     curl_setopt($curl_handle, CURLOPT_POSTFIELDS, $fileContent);
//     $result = curl_exec($curl_handle);
//     $status = curl_getinfo($curl_handle, CURLINFO_HTTP_CODE);

//     if ($result === false) {
//         $status = curl_errno($curl_handle);
//         $result = 'put file to oss - curl error :' . curl_error($curl_handle);
//     }
//     curl_close($curl_handle);
// //    $this->debug($url, $fileContent, $header, $result);
//     return $status;
// }









