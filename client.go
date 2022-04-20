package times

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Slack struct {
	BaseURL string
	token   string
	channel string
}

type SOp func(*Slack) error

func NewSlack(ops ...SOp) (*Slack, error) {
	s := &Slack{
		BaseURL: "https://slack.com/api",
	}
	for _, op := range ops {
		if err := op(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func fillAuth(token, channel string) func(s *Slack) error {
	return func(s *Slack) error {
		s.token = token
		s.channel = channel
		return nil
	}
}

func (s *Slack) Auth() error {
	fmt.Println("Auth")
	url := s.BaseURL + "/auth.test"
	resp, err := http.Get(url)
	fmt.Println(resp)
	fmt.Println(err)
	return nil
}

// type PostForm struct {
// 	Token   string `json:"token"`
// 	Channel string `json:"channel"`
// 	Text    string `json:"text"`
// }

func (s *Slack) PostMessage(args []string) error {
	// 水曜日はjsonで呼び出すのが目標！
	if len(args) < 3 {
		return fmt.Errorf("expect more than 3 argument, get %s", strconv.Itoa(len(args)))
	}

	// ここからはまだ未検証
	// pf := PostForm{
	// 	Token:   os.Getenv("SLACK_API_TOKEN"),
	// 	Channel: os.Getenv("TIMES"),
	// 	Text:    args[2],
	// }

	// body, err := json.Marshal(pf)
	// if err != nil {
	// 	return err
	// }
	// req, err := http.NewRequest(http.MethodPost, s.BaseURL+"/chat.postMessage", bytes.NewBuffer(body))
	// req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// cli := &http.Client{}
	// resp, err := cli.Do(req)
	// // resp, err := http.Post(s.BaseURL+"/chat.postMessage", "application/json; charset=utf-8", bytes.NewBuffer(body))
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()
	// fmt.Println("============")

	// res, _ := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	// jsonStr := string(res)
	// fmt.Println("Response: ", jsonStr)
	// return nil

	// 以下は正常に動く
	form := url.Values{}
	form.Add("token", os.Getenv("SLACK_API_TOKEN"))
	form.Add("channel", os.Getenv("TIMES"))
	form.Add("text", args[2])
	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest(
		http.MethodPost,
		s.BaseURL+"/chat.postMessage",
		body,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	sb := string(respBody)
	fmt.Println(sb)

	fmt.Println("succesfully post message!!!!!!!")
	return nil
}

func (s *Slack) History() []string {
	return []string{}
}

func (s *Slack) auth() error {
	form := url.Values{}
	form.Add("token", os.Getenv("SLACK_API_TOKEN"))
	body1 := strings.NewReader(form.Encode())

	req, err := http.NewRequest(
		http.MethodPost,
		"https://slack.com/api/auth.test",
		body1,
	)
	if err != nil {
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	sb := string(body)
	fmt.Println(sb)
	defer resp.Body.Close()
	return nil
}

func (s *Slack) channelId() string {
	return ""
}
