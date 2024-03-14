package igateModel

import "encoding/json"

type MessagePayload struct {
	Msg     string `json:"msg"`
	From    string `json:"from"`
	AppCode string `json:"appCode,omitempty"`
	To      string `json:"to"`
}

type MessageResponse struct {
	ID            int             `json:"id,omitempty"`
	MessageID     string          `json:"message_id"`
	Response      json.RawMessage `json:"response"`
	TimeInserted  string          `json:"time_inserted"`
	ApplicationID int             `json:"application_id"`
	UserID        int             `json:"user_id"`
}
