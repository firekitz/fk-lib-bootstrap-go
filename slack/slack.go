package slack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var SLACK_URL = "https://hooks.slack.com/services/"

const (
	SLACK_CHANNEL_DEV    = ""
	SLACK_CHANNEL_STAGE  = ""
	SLACK_CHANNEL_META   = ""
	SLACK_CHANNEL_PROD   = ""
	SLACK_CHANNEL_EMERG  = "T016WSLQEVB/B03M4KC49QX/AysvutQL8C6Dy10GsU3Nloxu"
	SLACK_CHANNEL_ALERT  = ""
	SLACK_CHANNEL_CRIT   = ""
	SLACK_CHANNEL_ERROR  = "T016WSLQEVB/B03MG9QJHV3/yvT9PNcs6sqzvRUdecrwQyUv"
	SLACK_CHANNEL_WARN   = ""
	SLACK_CHANNEL_NOTICE = ""
	SLACK_CHANNEL_INFO   = ""
	SLACK_CHANNEL_DEBUG  = ""
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Action struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Url   string `json:"url"`
	Style string `json:"style"`
}

type Attachment struct {
	Fallback     *string   `json:"fallback"`
	Color        *string   `json:"color"`
	PreText      *string   `json:"pretext"`
	AuthorName   *string   `json:"author_name"`
	AuthorLink   *string   `json:"author_link"`
	AuthorIcon   *string   `json:"author_icon"`
	Title        *string   `json:"title"`
	TitleLink    *string   `json:"title_link"`
	Text         *string   `json:"text"`
	ImageUrl     *string   `json:"image_url"`
	Fields       []*Field  `json:"fields"`
	Footer       *string   `json:"footer"`
	FooterIcon   *string   `json:"footer_icon"`
	Timestamp    *int64    `json:"ts"`
	MarkdownIn   *[]string `json:"mrkdwn_in"`
	Actions      []*Action `json:"actions"`
	CallbackID   *string   `json:"callback_id"`
	ThumbnailUrl *string   `json:"thumb_url"`
}

type Payload struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconUrl     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	UnfurlMedia bool         `json:"unfurl_media,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	return attachment
}

func (attachment *Attachment) AddAction(action Action) *Attachment {
	attachment.Actions = append(attachment.Actions, &action)
	return attachment
}

func (payload *Payload) AddAttachment(attachment Attachment) *Payload {
	payload.Attachments = append(payload.Attachments, attachment)
	return payload
}

func (payload *Payload) SendSlack(channel string) {
	client := &http.Client{}
	b, err := json.Marshal(payload)
	req, err := http.NewRequest("POST", SLACK_URL+channel, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	data := string(body)

	log.Println(data)
}
