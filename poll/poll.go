package poll

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/volio/go-common/log"
	"github.com/volio/go-telegram/config"
	"github.com/volio/go-telegram/model"
	"go.uber.org/zap"
)

type Poll interface {
	Start(ch chan *model.Update)
	Stop()
}

func NewPoll(cfg *config.BotConfig) Poll {
	client := newHttpClient(cfg)
	return &poll{
		client: client,
		key:    cfg.Key,
	}
}

func newHttpClient(cfg *config.BotConfig) *http.Client {
	client := http.Client{}

	if cfg.EnableProxy {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(cfg.Proxy)
		}

		client.Transport = &http.Transport{
			Proxy: proxy,
		}
	}

	return &client
}

type poll struct {
	client  *http.Client
	key     string
	stopped bool
	offset  int64
}

func (p *poll) Start(ch chan *model.Update) {
	p.stopped = false
	for {
		if p.stopped {
			return
		}
		updates, err := p.fetchUpdates()
		if err != nil {
			log.L().With(zap.Error(err), zap.Int64("offset", p.offset)).Error("fetch updates failed")
		}
		for _, update := range updates {
			ch <- update
			if update.UpdateID > p.offset {
				p.offset = update.UpdateID
			}
		}
		if len(updates) == 0 || err != nil {
			time.Sleep(1 * time.Second)
		}
	}
}

func (p *poll) Stop() {
	p.stopped = true
}

func (p *poll) fetchUpdates() ([]*model.Update, error) {
	values := url.Values{}
	values.Set("offset", strconv.FormatInt(p.offset+1, 10))
	u := url.URL{
		Scheme:   "https",
		Host:     "api.telegram.org",
		Path:     fmt.Sprintf("/bot%s/getUpdates", p.key),
		RawQuery: values.Encode(),
	}

	resp, err := p.client.Get(u.String())
	if err != nil {
		return nil, errors.WithMessagef(err, "get resp error")
	}
	defer func() {
		if e := resp.Body.Close(); e != nil {
			log.L().With(zap.Error(e)).Error("close resp body failed")
		}
	}()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessagef(err, "error in read resp body")
	}
	var response model.UpdateReply
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, errors.WithMessagef(err, "unmarshal resp body failed, data: %v", string(data))
	}
	return response.Result, nil
}
