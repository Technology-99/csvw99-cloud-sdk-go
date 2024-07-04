/*
*

	@author: taco
	@Date: 2023/10/27
	@Time: 10:29

*
*/
package sdk

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/h5"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type H5 struct {
	h5.H5RpcServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewArticle(RpcClientConf *zrpc.RpcClientConf) *H5 {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	articleClient := h5.NewH5RpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &H5{
		H5RpcServiceClient: articleClient,
		Num:                1,
	}
}

func (c *Sdk) InitArticle() *Sdk {
	c.H5 = NewArticle(c.Config.RpcClientConf)
	c.H5.Status = types.STATUS_READY
	c.H5.Retry += 1
	return c
}

func (c *Sdk) ArticleCheckStatus() *Sdk {
	if c.H5.Status != types.STATUS_READY {
		if c.H5.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitArticle()
		}
	}
	return c
}

func (c *Sdk) ArticleCreate(in *h5.CreateArticleReq) (*h5.CreateArticleResp, error) {
	res, err := c.H5.ArticleCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleUpdate(in *h5.UpdateArticleReq) (*h5.Response, error) {
	res, err := c.H5.ArticleUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleDelete(in *h5.DeleteReq) (*h5.Response, error) {
	res, err := c.H5.ArticleDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleDeleteIds(in *h5.DeleteIdsReq) (*h5.Response, error) {
	res, err := c.H5.ArticleDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleQuery(in *h5.QueryReq) (*h5.QueryArticleResp, error) {
	res, err := c.H5.ArticleQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleQueryRecommendList(in *h5.QueryReq) (*h5.QueryRecommendListResp, error) {
	res, err := c.H5.ArticleQueryRecommendList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleQueryList(in *h5.QueryArticleListReq) (*h5.QueryArticleListResp, error) {
	res, err := c.H5.ArticleQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleQueryListQueryIds(in *h5.QueryIdsReq) (*h5.QueryArticleListResp, error) {
	res, err := c.H5.ArticleQueryListQueryIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleUpdateStatus(in *h5.UpdateStatusReq) (*h5.Response, error) {
	res, err := c.H5.ArticleUpdateStatus(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleBindArticleLabels(in *h5.UpdateBindLabelsReq) (*h5.Response, error) {
	res, err := c.H5.ArticleBindArticleLabels(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleUnBindArticleLabels(in *h5.UpdateBindLabelsReq) (*h5.Response, error) {
	res, err := c.H5.ArticleUnBindArticleLabels(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentCreate(in *h5.CreateAttachmentReq) (*h5.CreateResponse, error) {
	res, err := c.H5.AttachmentCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentUpdate(in *h5.UpdateAttachmentReq) (*h5.Response, error) {
	res, err := c.H5.AttachmentUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentDelete(in *h5.DeleteReq) (*h5.Response, error) {
	res, err := c.H5.AttachmentDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentDeleteIds(in *h5.DeleteIdsReq) (*h5.Response, error) {
	res, err := c.H5.AttachmentDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentQuery(in *h5.QueryReq) (*h5.QueryAttachmentResp, error) {
	res, err := c.H5.AttachmentQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentQueryList(in *h5.QueryAttachmentListReq) (*h5.QueryAttachmentListResp, error) {
	res, err := c.H5.AttachmentQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttaFileCreate(in *h5.CreateAttaFileReq) (*h5.CreateResponse, error) {
	res, err := c.H5.AttaFileCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttaFileDelete(in *h5.DeleteReq) (*h5.Response, error) {
	res, err := c.H5.AttaFileDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttaFileDeleteIds(in *h5.DeleteIdsReq) (*h5.Response, error) {
	res, err := c.H5.AttaFileDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttaFileQuery(in *h5.QueryReq) (*h5.QueryAttaFileResp, error) {
	res, err := c.H5.AttaFileQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttaFileQueryList(in *h5.QueryAttaFileListReq) (*h5.QueryAttaFileListResp, error) {
	res, err := c.H5.AttaFileQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// note: 相关站点的增删改查
func (c *Sdk) RelatedSitesCreate(in *h5.CreateRelatedSitesReq) (*h5.CreateResponse, error) {
	res, err := c.H5.RelatedSitesCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesUpdate(in *h5.UpdateRelatedSitesReq) (*h5.Response, error) {
	res, err := c.H5.RelatedSitesUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesDelete(in *h5.DeleteReq) (*h5.Response, error) {
	res, err := c.H5.RelatedSitesDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesDeleteIds(in *h5.DeleteIdsReq) (*h5.Response, error) {
	res, err := c.H5.RelatedSitesDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesQuery(in *h5.QueryReq) (*h5.QueryRelatedSitesResp, error) {
	res, err := c.H5.RelatedSitesQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesQueryList(in *h5.QueryRelatedSitesListReq) (*h5.QueryRelatedSitesListResp, error) {
	res, err := c.H5.RelatedSitesQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesQueryIds(in *h5.QueryIdsReq) (*h5.QueryRelatedSitesListResp, error) {
	res, err := c.H5.RelatedSitesQueryIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// note: 标签增删改查
func (c *Sdk) LabelCreate(in *h5.CreateLabelReq) (*h5.CreateLabelResp, error) {
	res, err := c.H5.LabelCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelUpdate(in *h5.UpdateLabelReq) (*h5.Response, error) {
	res, err := c.H5.LabelUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelDelete(in *h5.DeleteReq) (*h5.Response, error) {
	res, err := c.H5.LabelDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelDeleteIds(in *h5.DeleteIdsReq) (*h5.Response, error) {
	res, err := c.H5.LabelDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelQuery(in *h5.QueryReq) (*h5.QueryLabelResp, error) {
	res, err := c.H5.LabelQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelQueryList(in *h5.QueryLabelListReq) (*h5.QueryLabelListResp, error) {
	res, err := c.H5.LabelQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelUpdateStatus(in *h5.UpdateStatusReq) (*h5.Response, error) {
	res, err := c.H5.LabelUpdateStatus(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
