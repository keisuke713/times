package times

import (
	_ "fmt"
	"io"
)

type CmdName string

const (
	POST    CmdName = "post"
	HISTORY         = "history"
	HELP            = "help"
)

type Cmd interface {
	Name() string
	Usage() string
	MaxArg() int
	Run(io.Writer, []string) error
}

var CmdMap = map[CmdName]Cmd{
	POST: &PostCmd{},
	HELP: &HelpCmd{},
}
