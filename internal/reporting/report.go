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

type ReportDTO struct {
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func ConstructReportDTOFromStr(str string) ReportDTO {
	bytes := []byte(str)
	report := ReportDTO{}
	err := json.Unmarshal(bytes, &report)
	if err != nil {
		fmt.Println("error occurred when unmarshalling.")
		return ReportDTO{}
	}
	return report
}

func (r ReportDTO) IsMessageSentTo(receiver string) bool {
	return r.Receiver == receiver
}
