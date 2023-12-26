/*
*

	@author: taco
	@Date: 2023/10/27
	@Time: 10:40

*
*/
package main

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk"
	"os"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitEms().InitOss()
	s = s.AutoAuth().InitArticle()
	s.ArticleCheckStatus()
	select {}
}
