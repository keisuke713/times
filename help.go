package times

import (
	"fmt"
	"io"
	"sort"
	"text/tabwriter"
)

type HelpCmd struct{}

func (p *HelpCmd) Name() string {
	return "Help"
}

func (p *HelpCmd) Usage() string {
	return "shows usage"
}

func (p *HelpCmd) MaxArg() int {
	return 0
}

var desc = `
Usage: Times is incredibly cli tool that enable us to post message to times channel

SubCommands:
`

func ShowUsage(w io.Writer) error {
	cms := make([]string, len(CmdMap))
	i := 0
	for k := range CmdMap {
		cms[i] = string(k)
		i++
	}
	sort.Strings(cms)

	tw := tabwriter.NewWriter(w, 0, 4, 1, ' ', 0)
	fmt.Fprintf(tw, "%s", desc)
	for _, k := range cms {
		cn := CmdName(k)
		fmt.Fprintf(tw, "\t%s\t%s\n", CmdMap[cn].Name(), CmdMap[cn].Usage())
	}
	return tw.Flush()
}

func (p *HelpCmd) Run(out io.Writer, args []string) error {
	return ShowUsage(out)
}
