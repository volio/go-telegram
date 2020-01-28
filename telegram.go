package telegram

import (
	"github.com/volio/go-common/graceful"
	"github.com/volio/go-telegram/bot"
	"github.com/volio/go-telegram/config"
	"github.com/volio/go-telegram/longpoll"
	"github.com/volio/go-telegram/model"
	"github.com/volio/go-telegram/update"
)

type Telegram interface {
	Start()
	Stop()
	OnUpdate(fn update.HandleFunc)
	Bot() bot.Bot
}

func NewTelegram(cfg config.BotConfig) Telegram {
	return &telegram{
		bot:        bot.NewBot(&cfg),
		poll:       longpoll.NewLongPoll(&cfg),
		dispatcher: update.NewDispatcher(),
		q:          make(chan interface{}),
	}
}

type telegram struct {
	bot        bot.Bot
	poll       longpoll.LongPoll
	dispatcher update.Dispatcher
	q          chan interface{}
}

func (t *telegram) OnUpdate(fn update.HandleFunc) {
	t.dispatcher.SetHandler(fn)
}

func (t *telegram) Bot() bot.Bot {
	return t.bot
}

func (t *telegram) Start() {
	ch := make(chan *model.Update, 100)
	graceful.Go(func() {
		t.dispatcher.Run(ch, t.q)
	})
	t.poll.Start(ch)
}

func (t *telegram) Stop() {
	t.poll.Stop()
	close(t.q)
}
