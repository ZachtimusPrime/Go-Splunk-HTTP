package splunk

import (
	"net/http"
	"testing"
)

func TestEventCollectorResponseError(t *testing.T) {
	invalidEventNumber := 2
	ackID := 12345

	testCases := []struct {
		name   string
		input  *EventCollectorResponse
		expect string
	}{
		{
			name:   "Response is nil",
			input:  nil,
			expect: "",
		}, {
			name: "All response attributes are set",
			input: &EventCollectorResponse{
				Text:               "An error",
				Code:               10,
				InvalidEventNumber: &invalidEventNumber,
				AckID:              &ackID,
			},
			expect: "An error (Code: 10, InvalidEventNumber: 2, AckID: 12345)",
		}, {
			name: "Some response attributes are set",
			input: &EventCollectorResponse{
				Text:  "An error",
				Code:  10,
				AckID: &ackID,
			},
			expect: "An error (Code: 10, AckID: 12345)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errStr := tc.input.Error()
			if errStr != tc.expect {
				t.Errorf("Expected %q, got %q", tc.expect, errStr)
			}
		})
	}
}

func TestStatusCodeHTTPCode(t *testing.T) {
	testCases := []struct {
		name       string
		input      StatusCode
		expectCode int
		expectErr  bool
	}{
		{
			name:       "Known status code",
			input:      IncorrectIndex,
			expectCode: http.StatusBadRequest,
			expectErr:  false,
		}, {
			name:       "Unknown status code",
			input:      StatusCode(100),
			expectCode: -1,
			expectErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, err := tc.input.HTTPCode()

			if !tc.expectErr && err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tc.expectErr && err == nil {
				t.Fatalf("Expected an error to occur")
			}

			if code != tc.expectCode {
				t.Errorf("Expected %d, got %d", tc.expectCode, code)
			}
		})
	}
}
