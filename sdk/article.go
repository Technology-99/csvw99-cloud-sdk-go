/*
*

	@author: taco
	@Date: 2023/10/27
	@Time: 10:29

*
*/
package sdk

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/article"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type Article struct {
	article.ArticleRpcServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewArticle(RpcClientConf *zrpc.RpcClientConf) *Article {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	articleClient := article.NewArticleRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Article{
		ArticleRpcServiceClient: articleClient,
		Num:                     1,
	}
}

func (c *Sdk) InitArticle() *Sdk {
	c.Article = NewArticle(c.Config.RpcClientConf)
	c.Article.Status = types.STATUS_READY
	c.Article.Retry += 1
	return c
}

func (c *Sdk) ArticleCheckStatus() *Sdk {
	if c.Article.Status != types.STATUS_READY {
		if c.Article.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitArticle()
		}
	}
	return c
}

func (c *Sdk) ArticleCreate(in *article.CreateArticleReq) (*article.CreateArticleResp, error) {
	res, err := c.Article.Create(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleUpdate(in *article.UpdateArticleReq) (*article.Response, error) {
	res, err := c.Article.Update(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleDelete(in *article.DeleteReq) (*article.Response, error) {
	res, err := c.Article.Delete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleDeleteIds(in *article.DeleteIdsReq) (*article.Response, error) {
	res, err := c.Article.DeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleQuery(in *article.QueryReq) (*article.QueryArticleResp, error) {
	res, err := c.Article.Query(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleQueryRecommendList(in *article.QueryReq) (*article.QueryRecommendListResp, error) {
	res, err := c.Article.QueryRecommendList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleQueryList(in *article.QueryArticleListReq) (*article.QueryArticleListResp, error) {
	res, err := c.Article.QueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleUpdateStatus(in *article.UpdateStatusReq) (*article.Response, error) {
	res, err := c.Article.UpdateStatus(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleBindArticleLabels(in *article.UpdateBindLabelsReq) (*article.Response, error) {
	res, err := c.Article.BindArticleLabels(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleUnBindArticleLabels(in *article.UpdateBindLabelsReq) (*article.Response, error) {
	res, err := c.Article.UnBindArticleLabels(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentCreate(in *article.CreateAttachmentReq) (*article.CreateResponse, error) {
	res, err := c.Article.AttachmentCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentUpdate(in *article.UpdateAttachmentReq) (*article.Response, error) {
	res, err := c.Article.AttachmentUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentDelete(in *article.DeleteReq) (*article.Response, error) {
	res, err := c.Article.AttachmentDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentDeleteIds(in *article.DeleteIdsReq) (*article.Response, error) {
	res, err := c.Article.AttachmentDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentQuery(in *article.QueryReq) (*article.QueryAttachmentResp, error) {
	res, err := c.Article.AttachmentQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttachmentQueryList(in *article.QueryAttachmentListReq) (*article.QueryAttachmentListResp, error) {
	res, err := c.Article.AttachmentQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttaFileCreate(in *article.CreateAttaFileReq) (*article.CreateResponse, error) {
	res, err := c.Article.AttaFileCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttaFileDelete(in *article.DeleteReq) (*article.Response, error) {
	res, err := c.Article.AttaFileDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttaFileDeleteIds(in *article.DeleteIdsReq) (*article.Response, error) {
	res, err := c.Article.AttaFileDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttaFileQuery(in *article.QueryReq) (*article.QueryAttaFileResp, error) {
	res, err := c.Article.AttaFileQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) ArticleAttaFileQueryList(in *article.QueryAttaFileListReq) (*article.QueryAttaFileListResp, error) {
	res, err := c.Article.AttaFileQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// note: 相关站点的增删改查
func (c *Sdk) RelatedSitesCreate(in *article.CreateRelatedSitesReq) (*article.CreateResponse, error) {
	res, err := c.Article.RelatedSitesCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesUpdate(in *article.UpdateRelatedSitesReq) (*article.Response, error) {
	res, err := c.Article.RelatedSitesUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesDelete(in *article.DeleteReq) (*article.Response, error) {
	res, err := c.Article.RelatedSitesDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesDeleteIds(in *article.DeleteIdsReq) (*article.Response, error) {
	res, err := c.Article.RelatedSitesDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesQuery(in *article.QueryReq) (*article.QueryRelatedSitesResp, error) {
	res, err := c.Article.RelatedSitesQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesQueryList(in *article.QueryRelatedSitesListReq) (*article.QueryRelatedSitesListResp, error) {
	res, err := c.Article.RelatedSitesQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) RelatedSitesQueryIds(in *article.QueryIdsReq) (*article.QueryRelatedSitesListResp, error) {
	res, err := c.Article.RelatedSitesQueryIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
