package main

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/Technology-99/third_party/sony"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"time"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitSls()

	// note: 以下函数可以检查cloudc是否初始化成功，不调用也可以

	for i := 0; i < 1000; i++ {
		log := s.NewSlsLogMsg(&sdk.TemplateLog{
			AppName:     "news",
			AppPlatform: "web",
			RequestId:   sony.NextId(),
			Time:        time.Now().Format(types.TimeLayout),
			Path:        "/api/v1/mark",
			Method:      "GET",
			Body:        "{\"title\":\"a\",\"value\":\"b\",\"extra\":\"asd\"}",
			Message:     "test",
			UserId:      "1",
			UniqueId:    sony.NextId(),
		})

		if err := s.SlsCheckStatus().SlsSendLog(log, "192.168.1.1"); err != nil {
			panic(err)
		}
		logx.Infof("发送成功")
	}

	select {}
}
