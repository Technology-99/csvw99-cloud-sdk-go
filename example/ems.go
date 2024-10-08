package main

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/mix"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitMix()

	// note: 以下函数可以检查cloudc是否初始化成功，不调用也可以
	s = s.MixCheckStatus()

	res, err := s.MixCheckStatus().MixSendEms("", "测试邮件", "text/plain", "世界多么美好", []string{"xxxxx"}, []*mix.Cc{
		{
			Email: "xxxx",
			Name:  "xxxx",
		},
	})
	if err != nil {
		return
	}
	logx.Infof("打印一下请求的结果:%+v", res)
}
