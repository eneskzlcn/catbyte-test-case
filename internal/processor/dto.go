package processor

type MessageDTO struct {
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}
