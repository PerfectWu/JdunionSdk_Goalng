package Demo

import "JdunionSdk_Goalng/JdunionSdk"

var NewJDSdk JdunionSdk.JDSDKAPI
var AppKey string = ""    //京东联盟申请的应用 AppKey
var AppSecret string = "" //京东联盟申请的应用 AppSecret
func init() {
	JdunionSdk.New(AppKey, AppSecret)
	NewJDSdk = JdunionSdk.JDSDKConfig
}
