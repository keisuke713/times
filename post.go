package times

import (
	"fmt"
	"io"
	"strings"
)

const (
	BREAK = "\n"
)

type PostCmd struct{}

func (p *PostCmd) Name() string {
	return "post"
}

func (p *PostCmd) Usage() string {
	return "Post message to times channel. you must pass at least one message."
}

func (p *PostCmd) MaxArg() int {
	return 2
}

func (p *PostCmd) Run(out io.Writer, args []string) error {
	cli, err := NewSlack()
	if err != nil {
		return err
	}

	if err := cli.PostMessage(args); err != nil {
		return err
	}

	return nil
}

func (p *PostCmd) Example() string {
	return "`slack-times post message1 message2 as much as you want`"
}

type MessageForm struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func NewMessageForm(channel string, args []string) (*MessageForm, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("expect more than 1 argument, get 0")
	}

	return &MessageForm{
		Channel: channel,
		Text:    strings.Join(args, BREAK),
	}, nil
}
