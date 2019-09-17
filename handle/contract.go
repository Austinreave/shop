package handle
import (
	"shop/utils"
	"net/http"
	// "github.com/julienschmidt/httprouter"
	// "strings"
)

func AccountCreatePerson(r *http.Request, param []interface{}) ([]map[string]interface{}, error) {
	data := make(map[string]interface{})

	data["name"] = "房雨雨"
	data["idNo"] = "130726199205122232"
	data["idType"] = "19"

	utils.HttpPostData(data)
	return nil ,nil
}
