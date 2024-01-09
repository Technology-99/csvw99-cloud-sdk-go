package main

import (
	"os"

	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/douyin"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitDouyin()

	// note: 以下函数可以检查cloudc是否初始化成功，不调用也可以
	s = s.DouyinCheckStatus()

	res, err := s.DouyinCheckStatus().DouyinCode2Token(&douyin.CodeReq{
		Key:  "dbl_game",
		Code: "hG30DBMyBcBLPdM-X0wuj--dCLRNftp8r66NPCC87MeGSm2H0BPY-gkO2zpQbE30nVDtNF37lL2HzFYlOGVvvDxsh9xQvbFlg5nPAtIMqLu6_mDlvRGn1otqbYg",
		// AnonymousCode: "",
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", res.Data)
}
