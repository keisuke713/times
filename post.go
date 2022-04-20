package times

import (
	"io"
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
