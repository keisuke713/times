package times

import (
	"encoding/json"
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
	Token   string
	Channel string
}

func NewSlack() (*Slack, error) {
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("not exist token")
	}

	channel := os.Getenv("TIMES")
	if channel == "" {
		return nil, fmt.Errorf("not exist channel")
	}

	s := &Slack{
		BaseURL: "https://slack.com/api",
		Token:   token,
		Channel: channel,
	}

	return s, nil
}

// type PostForm struct {
// 	Token   string `json:"token"`
// 	Channel string `json:"channel"`
// 	Text    string `json:"text"`
// }

func (s *Slack) PostMessage(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("expect more than 3 argument, get %s", strconv.Itoa(len(args)))
	}
	// todo return error if argument is more than 5

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
	form.Add("token", s.Token)
	form.Add("channel", s.Channel)
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
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error: status code", resp.StatusCode)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	var resBodyMap map[string]interface{}
	json.Unmarshal(respBody, &resBodyMap)
	if resBodyMap["error"] != nil {
		return fmt.Errorf("%s", resBodyMap["error"])
	}

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
