package JdunionSdk

import "time"

const (
	JD_HOST            string = "https://router.jd.com/api?"
	CUSTOMREQMETHODGET string = "GET"
)

//定义接口
type JDSDKAPI interface {
	ConversionLink(urls string) *SubunionidResult //链接转换
	GetJdGoods(UriQuery string) *GoodsResult
	GetCategoryList(UriQuery string) *CateGoryResult
	GetGoodsJFen(param string) *GoodsResult
	GetOrders(ParamJsons string) *OrderResult
	SetSignJointUrlParam(Method string, param string) *Jdsdk
}

type Param struct {
	Method       string `json:"method"`
	App_key      string `json:"app_key"`
	Access_token string `json:"access_token"`
	Timestamp    string `json:"timestamp"`
	Format       string `json:"format"`
	V            string `json:"v"`
	Sign_method  string `json:"sign_method"`
	Param_json   string `json:"param_json"`
}

type Jdsdk struct {
	RequestParam Param
	SignAndUri   string
	RequestRul   string
}

var JDSDKConfig *Jdsdk

var JdAppKey, JdAppSecret string

func New(AppId, AppSecret string) {
	JdAppKey = AppId
	JdAppSecret = AppSecret
	J := &Jdsdk{}
	J.RequestParam = Param{
		Method:      "",
		App_key:     JdAppKey,
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
		Format:      "json",
		Sign_method: "md5",
		V:           "1.0",
		Param_json:  "",
	}
	JDSDKConfig = J
}
