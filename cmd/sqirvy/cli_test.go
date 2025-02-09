package main

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func TestCliFlags(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		wantHelp        bool
		wantModel       string
		wantFunction    string
		wantFiles       []string
		wantTemperature int
		wantErr         bool
	}{
		{
			name:            "no flags",
			args:            []string{"app"},
			wantHelp:        false,
			wantModel:       "claude-3-5-sonnet-latest",
			wantFunction:    "query",
			wantFiles:       []string{},
			wantTemperature: 50,
			wantErr:         false,
		},
		{
			name:            "short help flag",
			args:            []string{"app", "-h"},
			wantHelp:        true,
			wantModel:       "claude-3-5-sonnet-latest",
			wantFunction:    "query",
			wantFiles:       []string{},
			wantTemperature: 50,
			wantErr:         false,
		},
		{
			name:            "short model flag",
			args:            []string{"app", "-m", "gpt-4"},
			wantHelp:        false,
			wantModel:       "gpt-4",
			wantFunction:    "query",
			wantFiles:       []string{},
			wantTemperature: 50,
			wantErr:         false,
		},
		{
			name:            "short function flag",
			args:            []string{"app", "-f", "query"},
			wantHelp:        false,
			wantModel:       "claude-3-5-sonnet-latest",
			wantFunction:    "query",
			wantFiles:       []string{},
			wantTemperature: 50,
			wantErr:         false,
		},
		{
			name:            "all flags",
			args:            []string{"app", "-h", "-m", "gpt-4", "-f", "query"},
			wantHelp:        true,
			wantModel:       "gpt-4",
			wantFunction:    "query",
			wantFiles:       []string{},
			wantTemperature: 50,
			wantErr:         false,
		},
		{
			name:            "single file after flags",
			args:            []string{"app", "-m", "gpt-4", "file1.txt"},
			wantHelp:        false,
			wantModel:       "gpt-4",
			wantFunction:    "query",
			wantFiles:       []string{"file1.txt"},
			wantTemperature: 50,
			wantErr:         false,
		},
		{
			name:            "multiple files after flags",
			args:            []string{"app", "-m", "gpt-4", "file1.txt", "file2.txt", "file3.txt"},
			wantHelp:        false,
			wantModel:       "gpt-4",
			wantFunction:    "query",
			wantFiles:       []string{"file1.txt", "file2.txt", "file3.txt"},
			wantTemperature: 50,
			wantErr:         false,
		},
		{
			name:            "all flags with files",
			args:            []string{"app", "-h", "-m", "gpt-4", "-f", "query", "file1.txt", "file2.txt"},
			wantHelp:        true,
			wantModel:       "gpt-4",
			wantFunction:    "query",
			wantFiles:       []string{"file1.txt", "file2.txt"},
			wantTemperature: 50,
			wantErr:         false,
		},
		{
			name:            "files with spaces in names",
			args:            []string{"app", "-m", "gpt-4", "my file.txt", "another file.md"},
			wantHelp:        false,
			wantModel:       "gpt-4",
			wantFunction:    "query",
			wantFiles:       []string{"my file.txt", "another file.md"},
			wantTemperature: 50,
			wantErr:         false,
		},
		{
			name:            "temperature flag",
			args:            []string{"app", "-t", "75"},
			wantHelp:        false,
			wantModel:       "claude-3-5-sonnet-latest",
			wantFunction:    "query",
			wantFiles:       []string{},
			wantTemperature: 75,
			wantErr:         false,
		},
		{
			name:            "temperature flag with other flags",
			args:            []string{"app", "-m", "gpt-4", "-t", "25", "-f", "query", "file1.txt"},
			wantHelp:        false,
			wantModel:       "gpt-4",
			wantFunction:    "query",
			wantFiles:       []string{"file1.txt"},
			wantTemperature: 25,
			wantErr:         false,
		},
		{
			name:            "invalid temperature flag (out of range)",
			args:            []string{"app", "-t", "150"},
			wantHelp:        false,
			wantModel:       "claude-3-5-sonnet-latest",
			wantFunction:    "query",
			wantFiles:       []string{},
			wantTemperature: 0, // Value doesn't matter in this case
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset command-line flags before each test
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
			os.Args = tt.args

			help, model, function, temperature, files, err := CliFlags()
			if (err != nil) != tt.wantErr {
				t.Errorf("CliFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if help != tt.wantHelp && !tt.wantErr {
				t.Errorf("CliFlags() help = %v, want %v", help, tt.wantHelp)
			}
			if model != tt.wantModel && !tt.wantErr {
				t.Errorf("CliFlags() model = %v, want %v", model, tt.wantModel)
			}
			if function != tt.wantFunction && !tt.wantErr {
				t.Errorf("CliFlags() function = %v, want %v", function, tt.wantFunction)
			}
			if !reflect.DeepEqual(files, tt.wantFiles) && !tt.wantErr {
				t.Errorf("CliFlags() files = %v, want %v", files, tt.wantFiles)
			}
			if temperature != tt.wantTemperature && !tt.wantErr {
				t.Errorf("CliFlags() temperature = %v, want %v", temperature, tt.wantTemperature)
			}
		})
	}
}
