package sdk

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/thirdparty"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type ThirdParty struct {
	thirdparty.ThirdPartyRpcServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewThirdParty(RpcClientConf *zrpc.RpcClientConf) *ThirdParty {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := thirdparty.NewThirdPartyRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &ThirdParty{
		ThirdPartyRpcServiceClient: client,
		Num:                        1,
	}
}

func (s *Sdk) InitThirdParty() *Sdk {
	s.ThirdParty = NewThirdParty(s.Config.RpcClientConf)
	s.ThirdParty.Status = types.STATUS_READY
	s.ThirdParty.Retry += 1
	return s
}

func (s *Sdk) ThirdPartyCheckStatus() *Sdk {
	if s.ThirdParty.Status != types.STATUS_READY {
		if s.ThirdParty.Retry >= s.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			s.InitThirdParty()
		}
	}
	return s
}

func (s *Sdk) WechatWebRedirectWechat(in *thirdparty.WechatWebRedirectReq) (*thirdparty.WechatWebRedirectResp, error) {
	res, err := s.ThirdParty.WechatWebRedirectWechat(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) WechatCode2Token(in *thirdparty.WechatCodeReq) (*thirdparty.WechatCodeResp, error) {
	res, err := s.ThirdParty.WechatCode2Token(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) WechatMiniAppCode2Session(in *thirdparty.WechatCodeReq) (*thirdparty.WechatMiniAppCodeResp, error) {
	res, err := s.ThirdParty.WechatMiniAppCode2Session(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) WechatMiniAppCode2Phone(in *thirdparty.WechatCodeReq) (*thirdparty.WechatMiniAppCode2PhoneResp, error) {
	res, err := s.ThirdParty.WechatMiniAppCode2Phone(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) WechatRefreshUserToken(in *thirdparty.WechatRefreshReq) (*thirdparty.WechatRefreshResp, error) {
	res, err := s.ThirdParty.WechatRefreshUserToken(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) WechatUserToken2UserInfo(in *thirdparty.WechatTokenReq) (*thirdparty.WechatUserInfoResp, error) {
	res, err := s.ThirdParty.WechatUserToken2UserInfo(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) WechatOfficialAccountAccessToken(in *thirdparty.WechatOaKeyReq) (*thirdparty.WechatOaAccessTokenResp, error) {
	res, err := s.ThirdParty.WechatOfficialAccountAccessToken(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) WechatWebAutoRedirectWechat(in *thirdparty.WechatWebAutoRedirectReq) (*thirdparty.WechatWebAutoRedirectResp, error) {
	res, err := s.ThirdParty.WechatWebAutoRedirectWechat(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) WechatOfficialAccountJsApiTicket(in *thirdparty.WechatOaKeyReq) (*thirdparty.WechatOaJsApiTicketResp, error) {
	res, err := s.ThirdParty.WechatOfficialAccountJsApiTicket(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) DYMiniGameCode2Token(in *thirdparty.DYMiniGameCode2TokenReq) (*thirdparty.DYMiniGameCode2TokenResp, error) {
	res, err := c.ThirdParty.DYMiniGameCode2Token(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) DYMiniGameOAAccessToken(in *thirdparty.DYMiniGameOAAccessTokenReq) (*thirdparty.DYMiniGameOAAccessTokenResp, error) {
	res, err := c.ThirdParty.DYMiniGameOAAccessToken(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
