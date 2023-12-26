/*
*

	@author: taco
	@Date: 2023/10/27
	@Time: 10:29

*
*/
package sdk

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/label"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type Label struct {
	label.LabelRpcServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewLabel(RpcClientConf *zrpc.RpcClientConf) *Label {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	labelClient := label.NewLabelRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Label{
		LabelRpcServiceClient: labelClient,
		Num:                   1,
	}
}

func (c *Sdk) InitLabel() *Sdk {
	c.Label = NewLabel(c.Config.RpcClientConf)
	c.Label.Status = types.STATUS_READY
	c.Label.Retry += 1
	return c
}

func (c *Sdk) LabelCheckStatus() *Sdk {
	if c.Label.Status != types.STATUS_READY {
		if c.Label.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitLabel()
		}
	}
	return c
}

func (c *Sdk) LabelCreate(in *label.CreateLabelReq) (*label.CreateLabelResp, error) {
	res, err := c.Label.Create(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelUpdate(in *label.UpdateLabelReq) (*label.Response, error) {
	res, err := c.Label.Update(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelDelete(in *label.DeleteReq) (*label.Response, error) {
	res, err := c.Label.Delete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelDeleteIds(in *label.DeleteIdsReq) (*label.Response, error) {
	res, err := c.Label.DeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelQuery(in *label.QueryReq) (*label.QueryLabelResp, error) {
	res, err := c.Label.Query(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelQueryList(in *label.QueryLabelListReq) (*label.QueryLabelListResp, error) {
	res, err := c.Label.QueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelUpdateStatus(in *label.UpdateStatusReq) (*label.Response, error) {
	res, err := c.Label.UpdateStatus(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
