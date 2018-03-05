package entity

// MailMessage represents income message struct
type MailMessage struct {
	Action  string                 `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}
