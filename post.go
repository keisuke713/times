package times

import (
	"io"
	"fmt"
	"strings"
)

const (
	BREAK = "\n"
)

type PostCmd struct{}

func (p *PostCmd) Name() string {
	return "Post"
}

func (p *PostCmd) Usage() string {
	return "posts message to times channel"
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

type MessageForm struct {
	Channel string `json:"channel"`
	Text string `json:"text"`
}

func NewMessageForm(channel string, args []string) (*MessageForm, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("expect more than 1 argument, get 0")
	}

	return &MessageForm{
		Channel: channel,
		Text: strings.Join(args, BREAK),
	}, nil
}