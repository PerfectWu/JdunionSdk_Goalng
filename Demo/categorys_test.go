package Demo

import (
	"testing"
)

func TestCutomInterfaceGoods(t *testing.T) {
	ParamJson := "{\"req\":{\"parentId\":0,\"grade\":0}}"
	categorys2 := NewJDSdkCutom.NewFunc(ParamJson)
	t.Log(categorys2)
}
func TestCateGorys(t *testing.T) {
	ParamJson := "{\"req\":{\"parentId\":0,\"grade\":0}}"
	categorys := NewJDSdk.GetCategoryList(ParamJson)
	t.Log(categorys)
}
