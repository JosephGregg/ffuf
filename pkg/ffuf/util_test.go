package ffuf

import (
	"math/rand"
	"testing"
)

func TestRandomString(t *testing.T) {
	length := 1 + rand.Intn(65535)
	str := RandomString(length)

	if len(str) != length {
		t.Errorf("Length of generated string was %d, was expecting %d", len(str), length)
	}
}

func TestUniqStringSlice(t *testing.T) {
	slice := []string{"foo", "foo", "bar", "baz", "baz", "foo", "baz", "baz", "foo"}
	expectedLength := 3

	uniqSlice := UniqStringSlice(slice)

	if len(uniqSlice) != expectedLength {
		t.Errorf("Length of slice was %d, was expecting %d", len(uniqSlice), expectedLength)
	}
}

func TestHostURLFromRequest(t *testing.T) {
	tests := []struct {
		name     string
		req      Request
		expected string
	}{
		{
			name: "Normal URL with path",
			req: Request{
				Host: "example.com",
				Url:  "http://example.com/path/to/file",
			},
			expected: "example.com/path/to",
		},
		{
			name: "URL with no path",
			req: Request{
				Host: "example.com",
				Url:  "http://example.com",
			},
			expected: "example.com",
		},
		{
			name: "URL with root path",
			req: Request{
				Host: "example.com",
				Url:  "http://example.com/",
			},
			expected: "example.com",
		},
		{
			name: "URL with single level path",
			req: Request{
				Host: "example.com",
				Url:  "http://example.com/file",
			},
			expected: "example.com",
		},
		{
			name: "URL with invalid format",
			req: Request{
				Host: "example.com",
				Url:  "invalid-url",
			},
			expected: "example.com",
		},
		{
			name: "URL with empty string",
			req: Request{
				Host: "example.com",
				Url:  "",
			},
			expected: "example.com",
		},
		{
			name: "URL with FireProx path",
			req: Request{
				Host: "api-id.execute-api.region.amazonaws.com",
				Url:  "https://api-id.execute-api.region.amazonaws.com/FUZZ",
			},
			expected: "api-id.execute-api.region.amazonaws.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HostURLFromRequest(tt.req)
			if result != tt.expected {
				t.Errorf("HostURLFromRequest() = %v, want %v", result, tt.expected)
			}
		})
	}
}
