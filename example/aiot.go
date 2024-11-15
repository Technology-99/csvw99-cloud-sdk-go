package main

import (
	"context"
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/aiot"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk"
	"os"
	"time"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC()

	// note: 以下函数可以检查cloudc是否初始化成功，不调用也可以
	s = s.CloudCCheckStatus()

	s.AIot.JStyleDeviceSleepParse(context.Background(), &aiot.JStyleDeviceSleepParseReq{
		Key:          "test",
		SyncDateTime: time.Now().UnixMilli(),
		Data:         "",
	})

}
