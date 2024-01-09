package sdk

import (
	"errors"

	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/cloudc"
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/douyin"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/Technology-99/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Douyin struct {
	douyin.DouyinRpcServiceClient
	Num         int64
	Status      int
	Retry       int
	Configs     map[string]*cloudc.ModelDouyinConfig
	UsingConfig *cloudc.ModelDouyinConfig
}

func NewDouyin(RpcClientConf *zrpc.RpcClientConf) *Douyin {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := douyin.NewDouyinRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Douyin{
		DouyinRpcServiceClient: client,
		Num:                    1,
	}
}

func (s *Douyin) GetAllConfigs() map[string]*cloudc.ModelDouyinConfig {
	return s.Configs
}

func (s *Douyin) GetConfig() (*cloudc.ModelDouyinConfig, error) {
	if s.UsingConfig != nil {
		return s.UsingConfig, nil
	}
	return nil, errors.New("not set config")
}

func (s *Douyin) SetConfig(key string) error {
	if _, ok := s.Configs[key]; !ok {
		return errors.New("not exist")
	}
	s.UsingConfig = s.Configs[key]
	return nil
}

func (c *Sdk) InitDouyin() *Sdk {
	c.Douyin = NewDouyin(c.Config.RpcClientConf)
	c.Douyin.Status = types.STATUS_READY
	c.Douyin.Retry += 1
	_, err := c.DouyinLoadCloudConfig()
	if err != nil {
		logx.Errorf("oss cloud config sync err: %v+", err)
		return c
	}
	return c
}

func (c *Sdk) DouyinCheckStatus() *Sdk {
	if c.Douyin.Status != types.STATUS_READY {
		if c.Douyin.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitDouyin()
		}
	}
	return c
}

func (c *Sdk) DouyinLoadCloudConfig() (*Sdk, error) {
	// note: 加载云端配置
	res, err := c.CloudCCheckStatus().CloudCDouyinConfigGetAll()
	if err != nil {
		logx.Errorf("douyin cloud config sync err: %v+", err)
		return c, err
	}
	if res.Code != response.SUCCESS {
		logx.Errorf("douyin cloud config sync err: %v+", res.Msg)
		return c, errors.New(res.Msg)
	}
	temp := make(map[string]*cloudc.ModelDouyinConfig)
	temp = res.Data.Configs
	// note: 判断是否存在default的配置文件，如果不存在，给予用户警告提示
	if _, ok := temp["default"]; !ok {
		logx.Alert("请注意，您的账号下没有default的配置文件，将会导致请求发送失败，发送前请SetConfig(key string)")
	}

	c.Douyin.Configs = temp
	c.Douyin.UsingConfig = temp["default"]
	return c, nil
}

func (c *Sdk) DouyinCode2Token(in *douyin.CodeReq) (*douyin.CodeResp, error) {
	// note: 读取当前使用配置
	if c.Douyin.UsingConfig == nil {
		c.Douyin.UsingConfig = c.Douyin.Configs["default"]
	}
	res, err := c.Douyin.Code2Token(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) DouyinOfficialAccountAccessToken(in *douyin.OaKeyReq) (*douyin.OaAccessTokenResp, error) {
	// note: 读取当前使用配置
	if c.Douyin.UsingConfig == nil {
		c.Douyin.UsingConfig = c.Douyin.Configs["default"]
	}
	res, err := c.Douyin.OfficialAccountAccessToken(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
