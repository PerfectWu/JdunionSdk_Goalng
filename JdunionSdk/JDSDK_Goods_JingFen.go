package JdunionSdk

import (
	"encoding/json"
	"strings"
)

/*
	@name jd.union.open.category.goods.get 商品类目查询
	@des 根据商品的父类目id查询子类目id信息，通常用获取各级类目对应关系，以便将推广商品归类。业务参数parentId、grade都输入0可查询所有一级类目ID，之后再用其作为parentId查询其子类目。
*/

func (J *Jdsdk) GetGoodsJFen(param string) (goodresult *GoodsResult) {
	Method := "jd.union.open.goods.jingfen.query"
	J.SetSignJointUrlParam(Method, param)
	var urls strings.Builder
	urls.WriteString(JD_HOST)
	urls.WriteString(J.SignAndUri)
	body, _ := HttpGet(urls.String())
	result := &Jd_union_open_goods_query_response{}
	e := json.Unmarshal([]byte(body), &result)
	if e != nil {
		panic(e)
	}
	goodresult = &GoodsResult{}
	e = json.Unmarshal([]byte(result.Jd_union_open_goods_query_response.Result), goodresult)
	if e != nil {
		panic(e)
	}
	return goodresult
}
