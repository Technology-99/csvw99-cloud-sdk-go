package main

import (
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

	logx.Infof("打印sdk版本号: %s", s.GetVersion())

	result, err := s.MixCheckStatus().MixSendSms("default", "13986537164", "520")
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下结果:%+v", result)
}
