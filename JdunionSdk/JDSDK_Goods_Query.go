package JdunionSdk

import (
	"encoding/json"
	"strings"
	"time"
)

type Jd_union_open_order_query_response struct {
	Jd_union_open_order_query_response struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_order_query_response"`
}

type OrderResult struct {
	Code      int    `json:"code"`
	HasMore   bool   `json:"hasMore"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
	Data      []struct {
		FinishTime int64     `json:"finishTime"`
		OrderEmt   int       `json:"orderEmt"`
		OrderId    int64     `json:"orderId"`
		OrderTime  time.Time `json:"orderTime"`
		ParentId   int64     `json:"parentId"`
		Plus       int       `json:"plus"`
		PopId      int       `json:"popId"`
		SkuList    []struct {
			Price       float64 `json:"price"`
			SkuName     string  `json:"skuName"`
			SubSideRate int     `json:"subSideRate"`
			SubsidyRate int     `json:"subsidyRate"`
		}
	}
}
type GoodsResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		BrandCode         string  `json:"brandCode"`
		BrandName         string  `json:"brandName"`
		GoodCommentsShare float64 `json:"goodCommentsShare"`
		CouponInfo        struct {
			CouponList []struct {
				Discount float64 `json:"discount"`
				Link     string  `json:"link"`
				IsBest   int     `json:"isBest"`
			} `json:"couponList"`
		} `json:"couponInfo"`
		ShopInfo struct {
			ShopId int `json:"shopId"`
		} `json:"shopInfo"`
		SkuId     int64  `json:"skuId"`
		SkuName   string `json:"skuName"`
		PriceInfo struct {
			Price       float64 `json:"price"`
			LowestPrice float64 `json:"lowestprice"`
		} `json:"priceInfo"`
		ImageInfo struct {
			ImageList []struct {
				Url string `json:"url"`
			} `json:"imageList"`
			IsJdSale int `json:"isJdSale"`
		} `json:"imageInfo"`
		Comments       int `json:"comments"`
		CommissionInfo struct {
			Commission      float64 `json:"commission"`
			CommissionShare float64 `json:"commissionshare"`
		} `json:"commissionInfo"`
	} `json:"data"`
	TotalCount int    `json:"totalCount"`
	RequestId  string `json:"requestId"`
}

//获取商品列表
type Jd_union_open_goods_query_response struct {
	Jd_union_open_goods_query_response struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_goods_query_response"`
}

//获取京东联盟的商品链接
func (J *Jdsdk) GetJdGoods(UriQuery string) (GoodResult *GoodsResult) {
	Method := "jd.union.open.goods.query"
	J.SetSignJointUrlParam(Method, UriQuery)
	var urls strings.Builder
	urls.WriteString(JD_HOST)
	urls.WriteString(J.SignAndUri)
	body, _ := HttpGet(urls.String())
	result := &Jd_union_open_goods_query_response{}
	e := json.Unmarshal([]byte(body), &result)
	if e != nil {
		panic(e)
	}
	GoodResult = &GoodsResult{}
	e = json.Unmarshal([]byte(result.Jd_union_open_goods_query_response.Result), GoodResult)
	if e != nil {
		panic(e)
	}
	return GoodResult
}
