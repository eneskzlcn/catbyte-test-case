package reporting

import (
	"encoding/json"
	"fmt"
)

type Report struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func ConstructReportFromStr(str string) Report {
	bytes := []byte(str)
	report := Report{}
	err := json.Unmarshal(bytes, &report)
	if err != nil {
		fmt.Println("error occurred when unmarshalling.")
		return Report{}
	}
	return report
}
func (r Report) IsMessageSentBetween(sender, receiver string) bool {
	return r.Sender == sender && r.Receiver == receiver
}
