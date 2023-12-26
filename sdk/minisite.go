/*
*

	@author: taco
	@Date: 2023/10/27
	@Time: 10:29

*
*/
package sdk

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/minisite"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type Minisite struct {
	minisite.MinisiteRpcServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewMinisite(RpcClientConf *zrpc.RpcClientConf) *Minisite {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	articleClient := minisite.NewMinisiteRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Minisite{
		MinisiteRpcServiceClient: articleClient,
		Num:                      1,
	}
}

func (c *Sdk) InitMinisite() *Sdk {
	c.Minisite = NewMinisite(c.Config.RpcClientConf)
	c.Minisite.Status = types.STATUS_READY
	c.Minisite.Retry += 1
	return c
}

func (c *Sdk) MinisiteCheckStatus() *Sdk {
	if c.Minisite.Status != types.STATUS_READY {
		if c.Minisite.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitMinisite()
		}
	}
	return c
}

func (c *Sdk) MinisiteCreate(in *minisite.CreateMinisiteReq) (*minisite.Response, error) {
	res, err := c.Minisite.MinisiteCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteUpdate(in *minisite.UpdateMinisiteReq) (*minisite.Response, error) {
	res, err := c.Minisite.MinisiteUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteDelete(in *minisite.DeleteReq) (*minisite.Response, error) {
	res, err := c.Minisite.MinisiteDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteDeleteIds(in *minisite.DeleteIdsReq) (*minisite.Response, error) {
	res, err := c.Minisite.MinisiteDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteQuery(in *minisite.QueryReq) (*minisite.QueryMinisiteResp, error) {
	res, err := c.Minisite.MinisiteQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteQueryList(in *minisite.QueryMinisiteListReq) (*minisite.QueryMinisiteListResp, error) {
	res, err := c.Minisite.MinisiteQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteUpdateStatus(in *minisite.UpdateStatusReq) (*minisite.Response, error) {
	res, err := c.Minisite.MinisiteUpdateStatus(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteUserCreate(in *minisite.CreateUserReq) (*minisite.Response, error) {
	res, err := c.Minisite.UserCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteUserUpdate(in *minisite.UpdateUserReq) (*minisite.Response, error) {
	res, err := c.Minisite.UserUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteUserDelete(in *minisite.DeleteReq) (*minisite.Response, error) {
	res, err := c.Minisite.UserDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteUserDeleteIds(in *minisite.DeleteIdsReq) (*minisite.Response, error) {
	res, err := c.Minisite.UserDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteUserQuery(in *minisite.QueryReq) (*minisite.QueryUserResp, error) {
	res, err := c.Minisite.UserQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteUserQueryList(in *minisite.QueryUserListReq) (*minisite.QueryUserListResp, error) {
	res, err := c.Minisite.UserQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) MinisiteUserUpdateStatus(in *minisite.UpdateStatusReq) (*minisite.Response, error) {
	res, err := c.Minisite.UserUpdateStatus(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
