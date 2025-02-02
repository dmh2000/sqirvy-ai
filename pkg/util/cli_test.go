package util

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func TestCliFlags(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantHelp     bool
		wantModel    string
		wantFunction string
		wantFiles    []string
		wantErr      bool
	}{
		{
			name:         "no flags",
			args:         []string{"app"},
			wantHelp:     false,
			wantModel:    "",
			wantFunction: "",
			wantFiles:    []string{},
			wantErr:      false,
		},
		{
			name:         "short help flag",
			args:         []string{"app", "-h"},
			wantHelp:     true,
			wantModel:    "",
			wantFunction: "",
			wantFiles:    []string{},
			wantErr:      false,
		},
		{
			name:         "short model flag",
			args:         []string{"app", "-m", "gpt-4"},
			wantHelp:     false,
			wantModel:    "gpt-4",
			wantFunction: "",
			wantFiles:    []string{},
			wantErr:      false,
		},
		{
			name:         "short function flag",
			args:         []string{"app", "-f", "analyze"},
			wantHelp:     false,
			wantModel:    "",
			wantFunction: "analyze",
			wantFiles:    []string{},
			wantErr:      false,
		},
		{
			name:         "all flags",
			args:         []string{"app", "-h", "-m", "gpt-4", "-f", "analyze"},
			wantHelp:     true,
			wantModel:    "gpt-4",
			wantFunction: "analyze",
			wantFiles:    []string{},
			wantErr:      false,
		},
		{
			name:         "single file after flags",
			args:         []string{"app", "-m", "gpt-4", "file1.txt"},
			wantHelp:     false,
			wantModel:    "gpt-4",
			wantFunction: "",
			wantFiles:    []string{"file1.txt"},
			wantErr:      false,
		},
		{
			name:         "multiple files after flags",
			args:         []string{"app", "-m", "gpt-4", "file1.txt", "file2.txt", "file3.txt"},
			wantHelp:     false,
			wantModel:    "gpt-4",
			wantFunction: "",
			wantFiles:    []string{"file1.txt", "file2.txt", "file3.txt"},
			wantErr:      false,
		},
		{
			name:         "all flags with files",
			args:         []string{"app", "-h", "-m", "gpt-4", "-f", "analyze", "file1.txt", "file2.txt"},
			wantHelp:     true,
			wantModel:    "gpt-4",
			wantFunction: "analyze",
			wantFiles:    []string{"file1.txt", "file2.txt"},
			wantErr:      false,
		},
		{
			name:         "files with spaces in names",
			args:         []string{"app", "-m", "gpt-4", "my file.txt", "another file.md"},
			wantHelp:     false,
			wantModel:    "gpt-4",
			wantFunction: "",
			wantFiles:    []string{"my file.txt", "another file.md"},
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset command-line flags before each test
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
			os.Args = tt.args

			help, model, function, files, err := CliFlags()
			if (err != nil) != tt.wantErr {
				t.Errorf("CliFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if help != tt.wantHelp {
				t.Errorf("CliFlags() help = %v, want %v", help, tt.wantHelp)
			}
			if model != tt.wantModel {
				t.Errorf("CliFlags() model = %v, want %v", model, tt.wantModel)
			}
			if function != tt.wantFunction {
				t.Errorf("CliFlags() function = %v, want %v", function, tt.wantFunction)
			}
			if !reflect.DeepEqual(files, tt.wantFiles) {
				t.Errorf("CliFlags() files = %v, want %v", files, tt.wantFiles)
			}
		})
	}
}
