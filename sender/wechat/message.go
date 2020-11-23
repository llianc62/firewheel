package wechat

// Message reflects the message body request WeChat api.
type Message struct {
	MessageType string           `json:"msgtype"`
	Text        *textMessage     `json:"text,omitempty"`
	Markdown    *markdownMessage `json:"markdown,omitempty"`
	Image       *imageMessage    `json:"image,omitempty"`
	News        *newsMessage     `json:"news,omitempty"`
	File        *fileMessage     `json:"file,omitempty"`
}

type textMessage struct {
	Content           string   `json:"content"`
	MentionList       []string `json:"mentioned_list,omitempty"`
	MentionMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

type markdownMessage struct {
	Content string `json:"content"`
}

type imageMessage struct {
	Base64 string `json:"base64"` // limitation: 2M
	MD5Sum string `json:"md5"`
}

type newsMessage struct {
	Articles []article `json:"articles"`
}

type article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

type fileMessage struct {
	MediaID string `json:"media_id"`
}
