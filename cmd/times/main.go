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
	// if err := run(os.Stdout, os.Stderr, []string{"times", "post", "nebashi"}); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%v \n", err)
	// 	os.Exit(ExitCodeParseFlagError)
	// }

	// test
	// form := url.Values{}
	// form.Add("token", os.Getenv("SLACK_API_TOKENN"))
	// // form.Add("token", os.Getenv("SLACK_API_TOKEN"))
	// form.Add("channel", os.Getenv("TIMES"))
	// form.Add("text", "takahashi")
	// body := strings.NewReader(form.Encode())

	// req, _ := http.NewRequest(
	// 	http.MethodPost,
	// 	"https://slack.com/api/chat.postMessage",
	// 	body,
	// )

	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// client := &http.Client{}
	// resp, _ := client.Do(req)

	// // エラーが見つかったらcustom errを返す
	// respBody, _ := ioutil.ReadAll(resp.Body)
	// defer resp.Body.Close()
	// sb := string(respBody)
	// fmt.Println(sb)
	// //これいけるっぽい
	// var test map[string]interface{}
	// json.Unmarshal(respBody, &test)
	// fmt.Println(test)

	// var data map[string]interface{}
	// json.NewDecoder(resp.Body).Decode(&data)
	// fmt.Println(data)

	// fmt.Println("succesfully post messages!!!!!!!")

	// todo
	// READMEをかけ、トークン取得の仕方（何のスコープが必要か）、環境変数のセット、各コマンドの説明

	if err := run(os.Stdout, os.Stderr, os.Args); err != nil {
		switch err := err.(type) {
		case *times.NotAuthedError:
			fmt.Fprintf(os.Stderr, "%v \n", err)
			fmt.Fprintf(os.Stderr, "%v \n", "You should get and set slack api token to environment variable whose name is SLACK_API_TOKEN.  \nIf you don't know how to do it, please check out https://github.com/keisuke713/times")
		case *times.ChannelNotFoundError:
			fmt.Fprintf(os.Stderr, "%v \n", err)
			fmt.Fprintf(os.Stderr, "%v \n", "You should set channel name to environment variable whose name is TIMES.  \n")
		default:
			fmt.Fprintf(os.Stderr, "%v \n", err)
		}
		os.Exit(ExitCodeParseFlagError)
	}
}

func run(stdout, stderr io.Writer, args []string) error {
	if len(args) < 2 {
		cmd := &times.HelpCmd{}
		if err := cmd.Run(stdout, []string{}); err != nil {
			return fmt.Errorf("help command is faild: %w", err)
		}
		return nil
	}

	sub := args[1]
	if cmd, ok := times.CmdMap[times.CmdName(sub)]; ok {
		if err := cmd.Run(stdout, args); err != nil {
			switch err := err.(type) {
			case *times.NotAuthedError:
				return err.AddPreText(fmt.Sprintf("%q command is failed: ", sub))
			case *times.ChannelNotFoundError:
				return err.AddPreText(fmt.Sprintf("%q command is failed: ", sub))
			default:
				return fmt.Errorf("%q command is failed: %q", sub, err)
			}
		}
	} else {
		return fmt.Errorf("unkown command %q", sub)
	}

	return nil
}
