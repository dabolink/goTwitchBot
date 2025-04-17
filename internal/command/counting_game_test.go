package command

import (
	"goWebsocket/internal/command/counter"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestCountingGameCommand_matches(t *testing.T) {
	type args struct {
		cmdInfo CommandInfo
	}
	tests := []struct {
		name string
		cmd  *CountingGameCommand
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cmd.matches(tt.args.cmdInfo); got != tt.want {
				t.Errorf("CountingGameCommand.matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountingGameCommand_parse(t *testing.T) {
	type args struct {
		cmdInfo CommandInfo
	}
	tests := []struct {
		name  string
		cmd   *CountingGameCommand
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.cmd.parse(tt.args.cmdInfo)
			if got != tt.want {
				t.Errorf("CountingGameCommand.parse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CountingGameCommand.parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCountingGameCommand_ID(t *testing.T) {
	tests := []struct {
		name string
		cmd  *CountingGameCommand
		want uuid.UUID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cmd.ID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountingGameCommand.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountingGameCommand_Process(t *testing.T) {
	tests := []struct {
		name    string
		cmd     *CountingGameCommand
		args    CommandInfo
		wantErr bool
	}{
		{
			name: "",
			cmd: &CountingGameCommand{
				counter: counter.NewCounter(),
			},
			args:    CommandInfo{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cmd.Process(tt.args)
			if tt.wantErr == (err == nil) {
				t.Errorf("unexpected error result %v", err)
			}
		})
	}
}
