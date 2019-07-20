package JdunionSdk

import (
	"encoding/json"
	"strings"
)

//获取订单
func (J *Jdsdk) GetOrders(ParamJsons string) (result *OrderResult) {
	Method := "jd.union.open.order.query "
	J.SetSignJointUrlParam(Method, ParamJsons)
	var urls strings.Builder
	urls.WriteString(JD_HOST)
	urls.WriteString(J.SignAndUri)
	body, _ := HttpGet(urls.String())
	response := &Jd_union_open_order_query_response{}
	e := json.Unmarshal([]byte(body), &response)
	if e != nil {
		panic(e)
	}
	result = &OrderResult{}
	e = json.Unmarshal([]byte(response.Jd_union_open_order_query_response.Result), result)
	if e != nil {
		panic(e)
	}
	return result
}
