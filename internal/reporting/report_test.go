package reporting_test

import (
	"github.com/eneskzlcn/catbyte-test-task/internal/reporting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructReportFromStr(t *testing.T) {
	str := "{\"receiver\":\"you\",\"message\":\"hi\"}"
	report := reporting.ConstructReportDTOFromStr(str)
	assert.Equal(t, report.Receiver, "you")
	assert.Equal(t, report.Message, "hi")
}
func TestReport_IsMessageSentBetween(t *testing.T) {
	type testCase struct {
		GivenReceiver string
		GivenReport   reporting.ReportDTO
		Expected      bool
	}
	testCases := []testCase{
		{
			GivenReceiver: "you",
			GivenReport: reporting.ReportDTO{
				Receiver: "you",
				Message:  "",
			},
			Expected: true,
		},
		{
			GivenReceiver: "you",
			GivenReport: reporting.ReportDTO{
				Receiver: "me",
				Message:  "",
			},
			Expected: false,
		},
	}
	for _, test := range testCases {
		result := test.GivenReport.IsMessageSentTo(test.GivenReceiver)
		assert.Equal(t, result, test.Expected)
	}
}
