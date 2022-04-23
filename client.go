package times

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Slack struct {
	BaseURL string
	Token   string
	Channel string
}

func NewSlack() (*Slack, error) {
	token := os.Getenv("SLACK_API_TOKEN")
	// token = ""
	if token == "" {
		return nil, &NotAuthedError{"not exist token"}
	}

	channel := os.Getenv("TIMES")
	// channel = ""
	if channel == "" {
		return nil, &ChannelNotFoundError{"not exist channel"}
	}

	s := &Slack{
		BaseURL: "https://slack.com/api",
		Token:   token,
		Channel: channel,
	}

	return s, nil
}

func (s *Slack) PostMessage(args []string) error {
	mf, err := NewMessageForm(s.Channel, args[2:])
	if err != nil {
		return err
	}

	body, err := json.Marshal(mf)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, s.BaseURL+"/chat.postMessage", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	bearer := "Bearer " + s.Token
	req.Header.Add("Authorization", bearer)

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

	timesId, err := s.timesId()
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("token", s.Token)
	form.Add("channel", timesId.Channel)
	form.Add("limit", hisotryCnt)
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

	for i := messagesLen - 1; 0 <= i; i-- {
		msg := messages.Messages[i]
		fmt.Println(msg.Text)
	}

	return nil
}

func (s *Slack) timesId() (*TimesId, error) {
	if TIMES_ID_CACHE[s.Channel] != nil {
		return TIMES_ID_CACHE[s.Channel], nil
	}

	form := url.Values{}
	form.Add("token", s.Token)
	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest(
		http.MethodPost,
		s.BaseURL+"/conversations.list",
		body,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error: status code", resp.StatusCode)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	var channels Channels
	json.Unmarshal(respBody, &channels)
	if channels.Error != "" {
		return nil, fmt.Errorf("%s", channels.Error)
	}

	timesId, err := channels.NewTimesId(s.Channel)
	if err == nil {
		TIMES_ID_CACHE[s.Channel] = timesId
	}
	return timesId, err
}

func (s *Slack) hisotryCnt(args []string) (string, error) {
	switch len(args) {
	case 2:
		return DEFAULT_HISTORY_CNT, nil
	case 3:
		return args[2], nil
	default:
		return "", fmt.Errorf("too much argument. must be less than 1 argument")
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
