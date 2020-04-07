package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/volio/go-common/log"
	"go.uber.org/zap"
)

func newHttpClient(timeout time.Duration, cfg *ProxyConfig) *http.Client {
	client := &http.Client{
		Timeout: timeout,
	}

	if cfg.Enable {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(cfg.Addr)
		}

		client.Transport = &http.Transport{
			Proxy: proxy,
		}
	}

	return client
}

func newBotClient(key string, c *http.Client) *botClient {
	return &botClient{
		hc:  c,
		key: key,
	}
}

type botClient struct {
	hc  *http.Client
	key string
}

func (c *botClient) DoPost(method string, v interface{}, r interface{}) error {
	u := url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   fmt.Sprintf("/bot%s/%s", c.key, method),
	}

	b, err := json.Marshal(v)
	if err != nil {
		return errors.WithMessage(err, "json marshal error")
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	if err != nil {
		return errors.WithMessage(err, "new request error")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return errors.WithMessagef(err, "get resp error")
	}
	defer func() {
		if e := resp.Body.Close(); e != nil {
			log.L().With(zap.Error(e)).Error("close resp body failed")
		}
	}()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithMessagef(err, "error in read resp body")
	}
	if err := json.Unmarshal(data, r); err != nil {
		return errors.WithMessagef(err, "unmarshal resp body failed, data: %v", string(data))
	}
	return nil
}

func (c *botClient) DoGet(method string, v map[string]string, r interface{}) error {
	values := url.Values{}
	for k, value := range v {
		values.Set(k, value)
	}
	u := url.URL{
		Scheme:   "https",
		Host:     "api.telegram.org",
		Path:     fmt.Sprintf("/bot%s/%s", c.key, method),
		RawQuery: values.Encode(),
	}

	resp, err := c.hc.Get(u.String())
	if err != nil {
		return errors.WithMessagef(err, "get resp error")
	}
	defer func() {
		if e := resp.Body.Close(); e != nil {
			log.L().With(zap.Error(e)).Error("close resp body failed")
		}
	}()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithMessagef(err, "error in read resp body")
	}
	if err := json.Unmarshal(data, r); err != nil {
		return errors.WithMessagef(err, "unmarshal resp body failed, data: %v", string(data))
	}
	return nil
}
