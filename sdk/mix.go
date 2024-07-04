package sdk

import (
	"fmt"
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/mix"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type Mix struct {
	mix.MixServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewMix(RpcClientConf *zrpc.RpcClientConf) *Mix {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := mix.NewMixServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Mix{
		MixServiceClient: client,
		Num:              1,
	}
}

func (s *Sdk) InitMix() *Sdk {
	s.Mix = NewMix(s.Config.RpcClientConf)
	s.Mix.Status = types.STATUS_READY
	s.Mix.Retry += 1
	return s
}

func (s *Sdk) MixCheckStatus() *Sdk {
	if s.Mix.Status != types.STATUS_READY {
		if s.Mix.Retry >= s.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			s.InitMix()
		}
	}
	return s
}

func (s *Sdk) MixCaptchaGenerateByCustom(DotCount, KeyLong, ImgWidth, ImgHeight int64, MaxSkew float64) (*mix.CaptchaResp, error) {
	res, err := s.Mix.CaptchaGenerate(s.SonyCtx(), &mix.CaptchaReq{
		DotCount:  DotCount,
		MaxSkew:   MaxSkew,
		KeyLong:   KeyLong,
		ImgWidth:  ImgWidth,
		ImgHeight: ImgHeight,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) MixCaptchaGenerateByKey(key string) (*mix.CaptchaResp, error) {
	res, err := s.Mix.CaptchaGenerate(s.SonyCtx(), &mix.CaptchaReq{
		Key: key,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) MixSendEms(key, Subject, SendType, SendBody string, RecipientEmail []string, cc []*mix.Cc) (*mix.EmsResp, error) {
	res, err := s.Mix.SendEmsRpc(s.SonyCtx(), &mix.EmsReq{
		Key:            key,
		RecipientEmail: RecipientEmail,
		Cc:             cc,
		Subject:        Subject,
		SendType:       SendType,
		SendBody:       SendBody,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) MixSendSms(key, phone string, args ...any) (*mix.SmsResp, error) {
	params := []string{}
	for _, arg := range args {
		params = append(params, fmt.Sprintf("%v", arg))
	}
	res, err := s.Mix.SendSms(s.SonyCtx(), &mix.SmsParams{
		Key:    key,
		Mobile: phone,
		Params: params,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) MixSendSmsAndKey(key, phone string, args ...any) (*mix.SmsResp, error) {
	params := []string{}
	for _, arg := range args {
		params = append(params, fmt.Sprintf("%v", arg))
	}
	res, err := s.Mix.SendSms(s.SonyCtx(), &mix.SmsParams{
		Key:    key,
		Mobile: phone,
		Params: params,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) MixGenerateUploadSign(key, Filename, UploadDir, CallBack string, IsCallBack bool) (*mix.GenerateUploadSignParamsResp, error) {
	res, err := s.Mix.GenerateUploadSign(s.SonyCtx(), &mix.GenerateUploadSignParams{
		Key:        key,
		Filename:   Filename,
		UploadDir:  UploadDir,
		IsCallBack: IsCallBack,
		CallBack:   CallBack,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
