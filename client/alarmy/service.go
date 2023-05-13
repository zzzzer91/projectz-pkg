package alarmy

import (
	"context"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
	"github.com/zzzzer91/httpgo"
)

type Service interface {
	GetHoroscope(ctx context.Context, zodiac, timezone string) (*GetHoroscopeResp, error)
}

func NewService() Service {
	return &serviceImpl{
		cli: httpgo.NewClient(10*time.Second, nil),
	}
}

type serviceImpl struct {
	cli *httpgo.Client
}

func (s *serviceImpl) GetHoroscope(ctx context.Context, zodiac, timezone string) (*GetHoroscopeResp, error) {
	url := fmt.Sprintf(urlHoroscope, zodiac, timezone)
	headers := []httpgo.Header{
		{Key: "User-Agent", Val: "Alarmy/7136 CFNetwork/1240.0.4 Darwin/20.6.0"},
		{Key: "Accept-Language", Val: "zh-cn"},
		{Key: "Accept", Val: "*/*"},
		// 自己不需要配置 Accept-Encoding 头部，原因如下：
		// if you manually set the Accept-Encoding request header,
		// than gzipped response will not automatically decompressed by the http.Transport.
		// Otherwise, behavior is controlled by the Transport's DisableCompression boolean
		// {"Accept-Encoding", "gzip, deflate, br"},
	}
	resp, err := s.cli.Get(ctx, url, headers...)
	if err != nil {
		return nil, errors.Wrap(err, "Get error")
	}
	defer resp.Body.Close()
	result := new(GetHoroscopeResp)
	if err := sonic.ConfigDefault.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, errors.Wrap(err, "Decode error")
	}
	return result, nil
}
