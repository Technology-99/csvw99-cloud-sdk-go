package sdk

import (
	"github.com/Technology-99/csvw99-cloud-sdk-go/sdk/types"
	"time"
)

type Config struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	Protocol        string

	AutoRetry        bool          `default:"false"`
	MaxRetryTimes    int           `default:"3"`
	Debug            bool          `default:"false"`
	Timeout          time.Duration `default:"5000"`
	AutoRefreshToken bool          `default:"true"`
	Deadline         int64
}

func DefaultConfig(AccessKeyId, AccessKeySecret string, Endpoint string) (config *Config) {
	config = &Config{
		AutoRetry:        true,
		MaxRetryTimes:    3,
		Debug:            false,
		Timeout:          defaultTimeout,
		AutoRefreshToken: true,
		AccessKeyId:      AccessKeyId,
		AccessKeySecret:  AccessKeySecret,
		Endpoint:         Endpoint,
		Deadline:         5,
		Protocol:         types.ProtocolHttps,
	}
	return
}

func NewConfig(c Config) *Config {
	return &c
}

func (c *Config) WithAutoRetry(isAutoRetry bool) *Config {
	c.AutoRetry = isAutoRetry
	return c
}

func (c *Config) WithMaxRetryTimes(MaxRetryTimes int) *Config {
	c.MaxRetryTimes = MaxRetryTimes
	return c
}

func (c *Config) WithDebug(Debug bool) *Config {
	c.Debug = Debug
	return c
}

func (c *Config) WithTimeout(Timeout time.Duration) *Config {
	c.Timeout = Timeout
	return c
}

func (c *Config) WithAutoRefreshToken(AutoRefreshToken bool) *Config {
	c.AutoRefreshToken = AutoRefreshToken
	return c
}

func (c *Config) WithAccessKeyId(AccessKeyId string) *Config {
	c.AccessKeyId = AccessKeyId
	return c
}

func (c *Config) WithAccessKeySecret(AccessKeySecret string) *Config {
	c.AccessKeySecret = AccessKeySecret
	return c
}

func (c *Config) WithDeadline(Deadline int64) *Config {
	c.Deadline = Deadline
	return c
}

func (c *Config) WithProtocol(Protocol string) *Config {
	c.Protocol = Protocol
	return c
}
