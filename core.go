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
	EXAMPLE         = "example"
)

type Cmd interface {
	Name() string
	Usage() string
	MaxArg() int
	Run(io.Writer, []string) error
	Example() string
}

var CmdMap = map[CmdName]Cmd{
	POST:    &PostCmd{},
	HISTORY: &HistoryCmd{},
	HELP:    &HelpCmd{},
	EXAMPLE: &ExampleCmd{},
}
