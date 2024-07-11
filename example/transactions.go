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

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitTransactions()

	logx.Infof("打印sdk版本号: %s", s.GetVersion())

	//for i := 1; i <= 39; i++ {
	//	logx.Infof("打印一下日志:%d", i)
	//	result, err := s.TransactionsCheckStatus().TransProductCreate(&transactions.ProductCreateReq{
	//		Key: "dbl_mini_app",
	//		Product: &transactions.ModelProduct{
	//			Name:    fmt.Sprintf("我能预知未来第%d集购买", i),
	//			Desc:    fmt.Sprintf("我能预知未来第%d集购买", i),
	//			Price:   1,
	//			PType:   4,
	//			Status:  2,
	//			Cover:   "https://dbl-prod.csvw88.com/movie/chongxinhuoyici/cover.jpg",
	//			Content: "https://dbl-prod.csvw88.com/movie/chongxinhuoyici/cover.jpg",
	//			Url:     "https://dbl-prod.csvw88.com/movie/chongxinhuoyici/cover.jpg",
	//			Stock:   -1,
	//			Extra:   "",
	//		},
	//	})
	//	if err != nil {
	//		panic(err)
	//	}
	//	logx.Infof("打印一下结果:%+v", result)
	//}

	//result2, err := s.TransactionsCheckStatus().TransCreateOrder(&transactions.CreateOrderParams{
	//	Key:           "dbl_mini_app",
	//	Platform:      "wechat",
	//	ProductSn:     "",
	//	Quantity:      0,
	//	ExtraData:     "",
	//	GoodsTags:     "",
	//	SupportFapiao: false,
	//	Openid:        "",
	//	SceneInfo:     nil,
	//	ProfitSharing: false,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//logx.Infof("打印一下结果:%+v", result2)
}
