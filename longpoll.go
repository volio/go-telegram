package telegram

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
	"go.uber.org/zap"
)

func NewLongPoll(key string, timeout time.Duration, client *http.Client, ch chan *Update) *longPoll {
	return &longPoll{
		client:  client,
		key:     key,
		ch:      ch,
		timeout: int64(timeout.Seconds()),
	}
}

type longPoll struct {
	client  *http.Client
	ch      chan *Update
	key     string
	stopped bool
	offset  int64
	timeout int64
}

func (p *longPoll) Start() {
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
			p.ch <- update
			if update.UpdateID > p.offset {
				p.offset = update.UpdateID
			}
		}
		if len(updates) == 0 || err != nil {
			time.Sleep(1 * time.Second)
		}
	}
}

func (p *longPoll) Stop() {
	p.stopped = true
}

func (p *longPoll) fetchUpdates() ([]*Update, error) {
	values := url.Values{}
	values.Set("offset", strconv.FormatInt(p.offset+1, 10))
	if p.timeout > 0 {
		values.Set("timeout", strconv.FormatInt(p.timeout, 10))
	}
	u := url.URL{
		Scheme:   "https",
		Host:     "api.telegram.org",
		Path:     fmt.Sprintf("/Bot%s/getUpdates", p.key),
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
	var r UpdateReply
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, errors.WithMessagef(err, "unmarshal resp body failed, data: %v", string(data))
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return r.Result, nil
}
