package checkversion

import (
	"testing"
)

func TestCheckVersion(t *testing.T) {
	tests := []struct {
		minVersion string
		userAgent  string
		shouldFail bool
	}{
		{"0.17.1", "/karlsend:0.17.1/karlsend:0.17.1/", false},
		{"0.17.1", "/karlsend:0.17.1/karlsend:0.0.0/", false},
		{"0.17.1", "/karlsend:0.17.1/karlsend:0.12.15(kdx_2.12.10)/", false},
		{"0.17.1", "/karlsend:0.18.9/karlsend:0.18.9/", false},
		{"0.17.1", "/karlsend:1.1.0/", false},

		{"0.18.9", "/karlsend:0.17.1/karlsend:0.17.1/", true},
		{"0.18.9", "/karlsend:0.17.1/karlsend:0.0.0/", true},
		{"0.18.9", "/karlsend:0.17.1/karlsend:0.12.15(kdx_2.12.10)/", true},
		{"0.18.9", "/karlsend:0.18.9/karlsend:0.18.9/", false},
		{"0.18.9", "/karlsend:1.1.0/", false},

		{"1.0.0", "/karlsend:0.17.1/karlsend:0.17.1/", true},
		{"1.0.0", "/karlsend:0.17.1/karlsend:0.0.0/", true},
		{"1.0.0", "/karlsend:0.17.1/karlsend:0.12.15(kdx_2.12.10)/", true},
		{"1.0.0", "/karlsend:0.18.9/karlsend:0.18.9/", true},
		{"1.0.0", "/karlsend:1.1.0/", false},
	}

	for _, tt := range tests {
		err := CheckVersion(tt.minVersion, tt.userAgent)
		if tt.shouldFail && err == nil {
			t.Errorf("Expected failure for %q with %q, but got nil", tt.minVersion, tt.userAgent)
		}
		if !tt.shouldFail && err != nil {
			t.Errorf("Unexpected error for %q with %q: %v", tt.minVersion, tt.userAgent, err)
		}
	}
}
