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
	// postは複数渡せるようにする
	// postにオプションつける、コード(コードだと隙間が出来るようねどうしようか)
	// 日付で区切る？
	// スレッド

	// 本番は必要
	if err := run(os.Stdout, os.Stderr, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v \n", err)
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
			return fmt.Errorf("%q command is failed: %q", sub, err)
		}
	} else {
		return fmt.Errorf("unkown command %q", sub)
	}

	return nil
}
