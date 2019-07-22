package Demo

import "JdunionSdk_Goalng/JdunionSdk"

var NewJDSdk JdunionSdk.JDSDKAPI

func init() {
	JdunionSdk.New(APPKEY, APPSECRET)
	NewJDSdk = JdunionSdk.JDSDKConfig
}
