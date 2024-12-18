package sdk

import (
	"context"
	"embed"
	"errors"
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"github.com/Technology-99/third_party/response"
	"github.com/Technology-99/third_party/sony"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
	"time"
)

const (
	keyRequestId = "requestId"
)

//go:embed "version"
var VersionF embed.FS

type Sdk struct {
	Version             string
	Config              *Config
	Status              int
	AccessToken         string
	AccessTokenExpires  int64
	RefreshToken        string
	RefreshTokenExpires int64
	RetryTimes          int
	Deadline            int64
	Context             context.Context
	Auth                *Auth
	CloudC              *CloudC
	ThirdParty          *ThirdParty
	Sls                 *Sls
	Mix                 *Mix
	H5                  *H5
	Transactions        *Transactions
	AIot                *AIot
}

func NewSdk() *Sdk {
	versionFile, err := VersionF.ReadFile("version")
	if err != nil {
		panic(err)
	}
	return &Sdk{
		Version:  string(versionFile),
		Status:   types.STATUS_NOT_READY,
		Context:  context.Background(),
		Deadline: 300,
	}
}
func (c *Sdk) GetVersion() string {
	return c.Version
}

func (c *Sdk) WithConfig(config *Config) *Sdk {
	c.Config = config
	return c
}

func (c *Sdk) WithDeadline(Deadline int64) *Sdk {
	c.Deadline = Deadline
	return c
}

func (c *Sdk) AutoAuth() *Sdk {
	c, err := c.InitAuth().AuthPing().AuthLogin()
	if err != nil {
		return nil
	}

	go c.AutoRefresh()

	return c
}

func (c *Sdk) InitAuth() *Sdk {
	c.Auth = NewAuth(c.Config.RpcClientConf)

	logx.Infof("ping types:%+v", c)

	return c
}

func (c *Sdk) AuthPing() *Sdk {
	ping, err := c.Auth.RequestPing()
	if err != nil {
		logx.Errorf("ping error:%v", err)
		return nil
	}

	if ping.Code == response.SUCCESS {
		c.Status = types.STATUS_READY
	} else {
		c.Status = types.STATUS_NOT_READY
	}

	return c
}

func (c *Sdk) AuthLogin() (*Sdk, error) {
	logx.Infof("打印sdk的状态 %s", c.Status)
	if c.Status == types.STATUS_NOT_READY {
		logx.Errorf("sdk not ready")
		return nil, types.ErrNotReady
	}
	res, err := c.Auth.RequestLogin(c.Config.AccessKeyId, c.Config.AccessKeySecret)
	if err != nil {
		logx.Errorf("ping error:%v", err)
		return nil, err
	}

	logx.Infof("RequestLogin types:%+v", res.String())

	if res.Code == response.SUCCESS {
		c.AuthSuccess()
		// note: 记录access token
		c.AccessToken = res.Data.AccessToken
		c.AccessTokenExpires = res.Data.AccessTokenExpires
		c.RefreshToken = res.Data.RefreshToken
		c.RefreshTokenExpires = res.Data.RefreshTokenExpires

	} else {
		c.AuthFail(errors.New(res.Msg))
		c, err = c.AuthLogin()
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Sdk) AuthRefresh() (*Sdk, error) {

	if c.RetryTimes >= c.Config.MaxRetryTimes {
		panic(types.ErrMaxErrTimes)
	}

	// note: 如果链接未准备好，直接返回
	if c.Status == types.STATUS_NOT_READY {
		return nil, types.ErrNotReady
	}

	// note: 判断accessToken过期了没
	if (c.AccessTokenExpires - c.Deadline) >= time.Now().Unix() {
		// note: 没过期，直接返回
		return c, nil
	}
	if (c.RefreshTokenExpires - c.Deadline) >= time.Now().Unix() {
		// note: refreshToken没过期，但是accessToken过期了
		res, err := c.Auth.RequestRefresh(c.Config.AccessKeyId, c.RefreshToken)
		if err != nil {
			return nil, err
		}
		if res.Code == response.SUCCESS {
			c.AuthSuccess()
			// note: 记录access token
			c.AccessToken = res.Data.AccessToken
			c.AccessTokenExpires = res.Data.AccessTokenExpires
		} else {
			c.AuthFail(errors.New(res.Msg))
			c, err = c.AuthRefresh()
			if err != nil {
				return nil, err
			}
			return c, nil
		}
	} else {
		// note: refreshToken过期了
		return c.AuthLogin()
	}

	return c, nil
}

func (c *Sdk) AutoRefresh() *Sdk {
	if c.Config.AutoRefreshToken {
		// note: check refresh token is expired
		for {
			if c.RetryTimes > c.Config.MaxRetryTimes {
				// note: close auto refresh
				c.Config.AutoRefreshToken = false
				logx.Errorf("RefreshToken fail: %+v", types.ErrMaxErrTimes)
				break
			}
			if _, err := c.AuthRefresh(); err != nil {
				logx.Errorf("RefreshToken fail: %+v", err)
				c.RetryTimes = c.RetryTimes - 1
				time.Sleep(time.Second)
				continue
				//return errs
			}
			time.Sleep(time.Second)
		}
	}
	return c
}

func (c *Sdk) AuthSuccess() {
	c.RetryTimes = 0
	c.Status = types.STATUS_LOGINED
}

func (c *Sdk) AuthFail(err error) {
	if c.Config.AutoRetry {
		c.RetryTimes++
		c.Status = types.STATUS_NOT_READY
	} else {
		panic(err)
	}
}

func (c *Sdk) WithContext(ctx context.Context) context.Context {
	c.Context = ctx
	return ctx
}

// note: 添加将requestID继承到下个服务的能力
func (c *Sdk) WithRequestId(requestId string) *Sdk {
	c.Context = context.WithValue(c.Context, keyRequestId, requestId)
	return c
}

func (c *Sdk) SonyCtx() context.Context {
	requestID := ""
	if value := c.Context.Value(keyRequestId); value != nil {
		requestID = value.(string)
	} else {
		requestID = sony.NextId()
	}
	logx.Debugf("requestId: %s", requestID)
	md := metadata.New(map[string]string{
		"X-RequestID-For": requestID,
		"X-AccessKey-For": c.Config.AccessKeyId,
		"Authorization":   "Bearer " + c.AccessToken,
	})
	c.Context = metadata.NewOutgoingContext(c.Context, md)
	return c.Context
}
