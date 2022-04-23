package times

import (
	"fmt"
	"io"
)

const (
	DEFAULT_HISTORY_CNT = "5"
)

var TIMES_ID_CACHE = map[string]*TimesId{}

type HistoryCmd struct{}

func (h *HistoryCmd) Name() string {
	return "History"
}

func (h *HistoryCmd) Usage() string {
	return "shows messages you posted"
}

func (h *HistoryCmd) MaxArg() int {
	return 1
}

func (h *HistoryCmd) Run(out io.Writer, args []string) error {
	cli, err := NewSlack()
	if err != nil {
		return err
	}

	if err := cli.History(args); err != nil {
		return err
	}

	return nil
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Channels struct {
	Error    string    `json:"error"`
	Channels []Channel `json:"channels"`
}

func (c *Channels) TimesId(name string) string {
	for _, ch := range c.Channels {
		if ch.Name == name {
			return ch.Id
		}
	}
	return ""
}

func (c *Channels) NewTimesId(name string) (*TimesId, error) {
	for _, ch := range c.Channels {
		if ch.Name == name {
			return &TimesId{
				Channel: ch.Id,
			}, nil
		}
	}
	return nil, fmt.Errorf("%s", "channel_not_found")
}

type Message struct {
	Text string `json:"text"`
}

type Messages struct {
	Error    string    `json:"error"`
	Messages []Message `json:"messages"`
}

type TimesId struct {
	Channel string `json:"channel"`
	Limit   string `json:"limit"`
}
