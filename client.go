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
	"bytes"
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

type MessageForm struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (s *Slack) PostMessage(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("expect more than 1 argument, get %s", strconv.Itoa(len(args)-2))
	}
	// todo return error if argument is more than 5
	if len(args) > 4 {
		return fmt.Errorf("too much argument. must be either 1 or 2 argument")
	}

	mf := MessageForm{
		Channel: s.Channel,
		Text:    args[2],
	}

	body, err := json.Marshal(mf)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, s.BaseURL+"/chat.postMessage", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	bearer := "Bearer " + s.Token
	req.Header.Add("authorization", bearer)

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

func (s *Slack) History(args []string) error {
	hisotryCnt, err := s.hisotryCnt(args)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	timesId, err := s.timedId()
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("token", s.Token)
	form.Add("channel", timesId)
	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest(
		http.MethodPost,
		s.BaseURL+"/conversations.history",
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
	var messages Messages
	json.Unmarshal(respBody, &messages)
	if messages.Error != "" {
		return fmt.Errorf("%s", messages.Error)
	}

	messagesLen := len(messages.Messages)

	start := hisotryCnt - 1
	if start >= messagesLen {
		start = messagesLen - 1
	}

	for i := start; 0<=i; i-- {
		msg := messages.Messages[i]
		fmt.Println(msg.Text)
	}

	return nil
}

func (s *Slack) timedId() (string, error) {
	form := url.Values{}
	form.Add("token", s.Token)
	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest(
		http.MethodPost,
		s.BaseURL+"/conversations.list",
		body,
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Error: status code", resp.StatusCode)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	var channels Channels
	json.Unmarshal(respBody, &channels)
	if channels.Error != "" {
		return "", fmt.Errorf("%s", channels.Error)
	}

	timedId := channels.TimesId(s.Channel)
	if timedId == "" {
		return "", fmt.Errorf("channel not found")
	}

	return timedId, nil
}

func (s *Slack) hisotryCnt(args []string) (int, error) {
	switch len(args) {
	case 2:
		return DEFAULT_HISTORY_CNT, nil
	case 3:
		return strconv.Atoi(args[2])
	default:
		return 0, fmt.Errorf("too much argument. must be 1 argument")
	}
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
