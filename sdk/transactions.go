package sdk

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/pb/transactions"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type Transactions struct {
	transactions.ProductRpcServiceClient
	transactions.TransactionsRpcServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewTransactions(RpcClientConf *zrpc.RpcClientConf) *Transactions {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client1 := transactions.NewProductRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())
	client2 := transactions.NewTransactionsRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())
	return &Transactions{
		ProductRpcServiceClient:      client1,
		TransactionsRpcServiceClient: client2,
		Num:                          1,
	}
}

func (s *Sdk) InitTransactions() *Sdk {
	s.Transactions = NewTransactions(s.Config.RpcClientConf)
	s.Transactions.Status = types.STATUS_READY
	s.Transactions.Retry += 1
	return s
}

func (s *Sdk) TransactionsCheckStatus() *Sdk {
	if s.Transactions.Status != types.STATUS_READY {
		if s.Transactions.Retry >= s.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			s.InitThirdParty()
		}
	}
	return s
}

func (s *Sdk) TransProductCreate(in *transactions.ProductCreateReq) (*transactions.ProductCreateResp, error) {
	res, err := s.Transactions.ProductCreate(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) TranssProductDelete(in *transactions.ProductDeleteReq) (*transactions.ProductDeleteResp, error) {
	res, err := s.Transactions.ProductDelete(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) TransProductDeleteIds(in *transactions.ProductDeleteIdsReq) (*transactions.ProductDeleteIdsResp, error) {
	res, err := s.Transactions.ProductDeleteIds(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) TransProductUpdate(in *transactions.ProductUpdateReq) (*transactions.ProductUpdateResp, error) {
	res, err := s.Transactions.ProductUpdate(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) TransProductQuery(in *transactions.ProductQueryReq) (*transactions.ProductQueryResp, error) {
	res, err := s.Transactions.ProductQuery(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) TransProductQueryIds(in *transactions.ProductQueryIdsReq) (*transactions.ProductQueryListResp, error) {
	res, err := s.Transactions.ProductQueryIds(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) TransCreateOrder(in *transactions.CreateOrderParams) (*transactions.CreateOrderResp, error) {
	res, err := s.Transactions.CreateOrder(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) TransQueryOrder(in *transactions.QueryOrderParams) (*transactions.QueryOrderResp, error) {
	res, err := s.Transactions.QueryOrder(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Sdk) TransCloseOrder(in *transactions.CloseOrderParams) (*transactions.CloseOrderResp, error) {
	res, err := s.Transactions.CloseOrder(s.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}