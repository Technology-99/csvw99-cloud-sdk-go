package main

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/cloudc"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"time"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitEms().InitOss()
	s.AutoAuth()
	for {
		res, err := s.CloudCCheckStatus().CloudCWechatConfigGet(&cloudc.ConfigGetParams{
			Key: "default",
		})
		if err != nil {
			return
		}
		logx.Infof("打印一下请求的结果:%+v", res)
		time.Sleep(time.Second * 5)
	}
}
