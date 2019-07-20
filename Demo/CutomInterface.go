package Demo

import (
	"JdunionSdk_Goalng/JdunionSdk"
	"encoding/json"
	"strings"
	"time"
)

type NewJDSDKAPI interface {
	JdunionSdk.JDSDKAPI
	NewFunc(s string) *JdunionSdk.CateGoryResult
}

type JdSdkCutom struct {
	JdunionSdk.Jdsdk
}

func (J *JdSdkCutom) NewFunc(UriQuery string) *JdunionSdk.CateGoryResult {
	Method := "jd.union.open.category.goods.get"
	J.SetSignJointUrlParam(Method, UriQuery)
	var urls strings.Builder
	urls.WriteString(JdunionSdk.JD_HOST)
	urls.WriteString(J.SignAndUri)
	body, _ := JdunionSdk.HttpGet(urls.String())
	result := &JdunionSdk.Reponse{}
	e := json.Unmarshal([]byte(body), &result)
	if e != nil {
		panic(e)
	}
	categoryResult := &JdunionSdk.CateGoryResult{}
	e = json.Unmarshal([]byte(result.Jd_union_open_category_goods_get_response.Result), categoryResult)
	if e != nil {
		panic(e)
	}
	return categoryResult
}

var JDSDKconfigCutom *JdSdkCutom

func New(AppKeyCutom, AppSecretCutom string) {

	JdunionSdk.JdAppKey = AppKeyCutom
	JdunionSdk.JdAppSecret = AppSecretCutom
	J := &JdSdkCutom{}
	J.RequestParam = JdunionSdk.Param{
		Method:      "",
		App_key:     JdunionSdk.JdAppKey,
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
		Format:      "json",
		Sign_method: "md5",
		V:           "1.0",
		Param_json:  "",
	}
	JDSDKconfigCutom = J
}

var AppKeyCutom string = ""    //京东联盟申请的应用 AppKey
var AppSecretCutom string = "" //京东联盟申请的应用 AppSecret
var NewJDSdkCutom NewJDSDKAPI

func init() {
	New(AppKeyCutom, AppSecretCutom)
	NewJDSdkCutom = JDSDKconfigCutom
}
