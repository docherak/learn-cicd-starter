package auth

import (
	"net/http"
	"testing"
)

// TestGetAPIKey checks various scenarios for retrieving an API key from HTTP headers.
func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectError bool
	}{
		{
			name:        "No Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectError: true,
		},
		{
			name: "Malformed Authorization Header",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			expectedKey: "",
			expectError: true,
		},
		{
			name: "Correct Authorization Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey 123456"},
			},
			expectedKey: "123456",
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)
			if err != nil && !tc.expectError {
				t.Errorf("GetAPIKey() unexpected error: %v", err)
			}
			if err == nil && tc.expectError {
				t.Errorf("GetAPIKey() expected error, got none")
			}
			if apiKey != tc.expectedKey {
				t.Errorf("GetAPIKey() got %v, want %v", apiKey, tc.expectedKey)
			}
		})
	}
}
