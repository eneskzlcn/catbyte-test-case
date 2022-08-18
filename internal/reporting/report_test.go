package reporting_test

import (
	"github.com/eneskzlcn/catbyte-test-task/internal/reporting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructReportFromStr(t *testing.T) {
	str := "{\"sender\":\"me\",\"receiver\":\"you\",\"message\":\"hi\"}"
	report := reporting.ConstructReportFromStr(str)
	assert.Equal(t, report.Sender, "me")
	assert.Equal(t, report.Receiver, "you")
	assert.Equal(t, report.Message, "hi")
}
func TestReport_IsMessageSentBetween(t *testing.T) {
	type testCase struct {
		GivenSender   string
		GivenReceiver string
		GivenReport   reporting.Report
		Expected      bool
	}
	testCases := []testCase{
		{
			GivenSender:   "me",
			GivenReceiver: "you",
			GivenReport: reporting.Report{
				Sender:   "me",
				Receiver: "you",
				Message:  "",
			},
			Expected: true,
		},
		{
			GivenSender:   "me",
			GivenReceiver: "you",
			GivenReport: reporting.Report{
				Sender:   "i",
				Receiver: "you",
				Message:  "",
			},
			Expected: false,
		},
	}
	for _, test := range testCases {
		result := test.GivenReport.IsMessageSentBetween(test.GivenSender, test.GivenReceiver)
		assert.Equal(t, result, test.Expected)
	}
}
