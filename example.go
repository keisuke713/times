package times

import (
	"fmt"
	"io"
	"sort"
	"text/tabwriter"
)

type ExampleCmd struct{}

func (e *ExampleCmd) Name() string {
	return "example"
}

func (e *ExampleCmd) Usage() string {
	return "Show example"
}

func (e *ExampleCmd) MaxArg() int {
	return 0
}

func ShowExample(w io.Writer) error {
	cms := make([]string, len(CmdMap))
	i := 0
	for k := range CmdMap {
		cms[i] = string(k)
		i++
	}
	sort.Strings(cms)

	tw := tabwriter.NewWriter(w, 0, 4, 1, ' ', 0)
	fmt.Fprintf(tw, "%s\n", "Example")
	for _, k := range cms {
		cn := CmdName(k)
		fmt.Fprintf(tw, "\t%s\t%s\n", CmdMap[cn].Name(), CmdMap[cn].Example())
	}
	return tw.Flush()
}

func (e *ExampleCmd) Run(out io.Writer, args []string) error {
	return ShowExample(out)
}

func (e *ExampleCmd) Example() string {
	return "`slack-times example`"
}
