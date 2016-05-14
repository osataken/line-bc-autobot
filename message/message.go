package message

type Receive struct {
	Result []Result `json:"result"`
}

type Result struct {
	From        string   `json:"from"`
	FromChannel int64   `json:"fromChannel"`
	To          []string `json:"to"`
	EventType   string   `json:"eventType"`
	Id          string   `json:"id"`
	Content     Content  `json:"content"`
}

type Content struct {
	Location        *string   `json:"location"`
	Id              string   `json:"id"`
	ContentType     int      `json:"contentType"`
	From            string   `json:"from"`
	CreatedTime     int64    `json:"createdTime"`
	To              []string `json:"to"`
	ToType          int      `json:"toType"`
	ContentMetadata map[string]string   `json:"contentMetadata"`
	Text            string   `json:"text"`
}

type EventsRequest struct {
	To        []string `json:"to,omitempty"`
	ToChannel int64    `json:"toChannel,omitempty"`
	EventType string   `json:"eventType,omitempty"`
	Content   *Content `json:"content,omitempty"`
}

type EventsResponse struct {
	Failed    []string `json:"failed,omitempty"`
	MessageId string   `json:"messageId,omitempty"`
	Timestamp int64    `json:"timestamp,omitempty"`
	version   int64    `json:"version,omitempty"`
}
