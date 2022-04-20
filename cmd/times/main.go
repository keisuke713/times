package main

import (
	"fmt"
	"io"
	"os"

	"github.com/keisuke713/times"
)

const (
	ExitCodeOK             = 0
	ExitCodeParseFlagError = 1
)

func main() {
	// res, _ := http.Post("https://slack.com/api/auth.test")
	// body, _ := ioutil.ReadAll((res.Body))
	// sb := string(body)
	// fmt.Println(sb)
	// fmt.Println("======")
	if err := run(os.Stdout, os.Stderr, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v \n", err)
		os.Exit(ExitCodeParseFlagError)
	}

}

func run(stdout, stderr io.Writer, args []string) error {
	if len(args) < 2 {
		// help command or custom error
		return nil
	}

	sub := args[1]
	if cmd, ok := times.CmdMap[times.CmdName(sub)]; ok {
		if err := cmd.Run(stdout, args); err != nil {
			// todo custom err
			return fmt.Errorf("%q command failed: %q", sub, err)
		}
	} else {
		// todo custom err
		return fmt.Errorf("unkown command %q", sub)
	}

	return nil
}
