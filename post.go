package times

import (
	"io"
)

type PostCmd struct{}

func (p *PostCmd) Name() string {
	return "post"
}

func (p *PostCmd) MaxArg() int {
	return 2
}

func (p *PostCmd) Usage() string {
	return "post message to times channel via terminal"
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
