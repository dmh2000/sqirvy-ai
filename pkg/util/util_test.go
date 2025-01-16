package util

import (
	"os"
	"strings"
	"testing"
)

// Test that ReadFile checks the max size
func TestReadFile(t *testing.T) {
	// create a large file
	largeFile, err := os.CreateTemp("", "large-file")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(largeFile.Name())

	_, err = largeFile.WriteString(strings.Repeat("a", 1024*1024*10)) // 10 MB
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		max     int64
		want    string
		want1   int64
		wantErr bool
	}{
		{
			name:    "10 MB",
			max:     1024 * 1024 * 10,
			want:    strings.Repeat("a", 1024*1024*10),
			want1:   1024 * 1024 * 10,
			wantErr: false,
		},
		{
			name:    "10 MB + 1",
			max:     1024*1024*10 - 1,
			want:    strings.Repeat("a", 1024*1024*10),
			want1:   1024 * 1024 * 10,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ReadFile(largeFile.Name(), tt.max)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			s := string(got)
			if s != tt.want {
				t.Errorf("ReadFile() got = %v, want %v", got, len(tt.want))
			}
			if got1 != tt.want1 {
				t.Errorf("ReadFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
