package sdk

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/aiot"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type AIot struct {
	aiot.AIotJStyleRpcServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewAIot(RpcClientConf *zrpc.RpcClientConf) *AIot {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client1 := aiot.NewAIotJStyleRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())
	return &AIot{
		AIotJStyleRpcServiceClient: client1,
		Num:                        1,
	}
}

func (s *Sdk) InitAIot() *Sdk {
	s.AIot = NewAIot(s.Config.RpcClientConf)
	s.AIot.Status = types.STATUS_READY
	s.AIot.Retry += 1
	return s
}

func (s *Sdk) AIotCheckStatus() *Sdk {
	if s.AIot.Status != types.STATUS_READY {
		if s.AIot.Retry >= s.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			s.InitThirdParty()
		}
	}
	return s
}

func (s *Sdk) AIotJStyleDeviceSleepParse(in *aiot.JStyleDeviceSleepParseReq) (*aiot.JStyleDeviceSleepParseResp, error) {
	res, err := s.AIot.JStyleDeviceSleepParse(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) AIotJStyleDeviceHeartParse(in *aiot.JStyleDeviceHeartParseReq) (*aiot.JStyleDeviceHeartParseResp, error) {
	res, err := s.AIot.JStyleDeviceHeartParse(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) AIotJStyleDeviceHeartOneParse(in *aiot.JStyleDeviceHeartOneParseReq) (*aiot.JStyleDeviceHeartOneParseResp, error) {
	res, err := s.AIot.JStyleDeviceHeartOneParse(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) AIotJStyleDeviceStepParse(in *aiot.JStyleDeviceStepParseReq) (*aiot.JStyleDeviceStepParseResp, error) {
	res, err := s.AIot.JStyleDeviceStepParse(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
