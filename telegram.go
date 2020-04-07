package telegram

import (
	"time"

	"github.com/volio/go-common/graceful"
)

func NewTelegram(cfg Config) *Telegram {
	ch := make(chan *Update, 100)
	phc := newHttpClient(time.Duration(0), &cfg.Proxy)

	return &Telegram{
		bot:        NewBot(&cfg),
		poll:       NewLongPoll(cfg.Bot.Key, cfg.Request.LongPollTimeout, phc, ch),
		dispatcher: newDispatcher(ch),
	}
}

type Telegram struct {
	bot        *Bot
	poll       *longPoll
	dispatcher *dispatcher
}

func (t *Telegram) OnUpdate(fn HandleFunc) {
	t.dispatcher.RegisterHandler(fn)
}

func (t *Telegram) Bot() *Bot {
	return t.bot
}

func (t *Telegram) Start() {
	graceful.Go(func() {
		t.dispatcher.Run()
	})
	t.poll.Start()
}

func (t *Telegram) Stop() {
	t.poll.Stop()
}
