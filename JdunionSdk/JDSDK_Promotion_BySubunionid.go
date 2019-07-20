package JdunionSdk

import (
	"encoding/json"
	"strings"
)

//转换链接
type Jd_union_open_promotion_bysubunionid_get_response struct {
	Jd_union_open_promotion_bysubunionid_get_response struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_promotion_bysubunionid_get_response"`
}

type SubunionidResult struct {
	Code int `json:"code"`
	Data struct {
		ClickURL string `json:"clickURL"`
		ShortURL string `json:"shortURL"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

//转换链接
func (J *Jdsdk) ConversionLink(Query string) (res *SubunionidResult) {
	Method := "jd.union.open.promotion.bysubunionid.get"
	J.SetSignJointUrlParam(Method, Query)
	var urls strings.Builder
	urls.WriteString(JD_HOST)
	urls.WriteString(J.SignAndUri)
	body, _ := HttpGet(urls.String())
	result := &Jd_union_open_promotion_bysubunionid_get_response{}
	e := json.Unmarshal([]byte(body), &result)
	if e != nil {
		panic(e)
	}
	res = &SubunionidResult{}
	e = json.Unmarshal([]byte(result.Jd_union_open_promotion_bysubunionid_get_response.Result), res)
	if e != nil {
		panic(e)
	}
	return res
}
