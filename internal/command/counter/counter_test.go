package counter

import (
	"reflect"
	"testing"
)

func TestCountManager_Increment(t *testing.T) {
	type args struct {
		key   string
		value int
	}
	tests := []struct {
		name    string
		cm      *Counter
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "first value - valid",
			cm:   NewCounter(),
			args: args{
				key:   "foo",
				value: 1,
			},
			wantErr: false,
		},
		{
			name: "not next number - error",
			cm:   NewCounter(),
			args: args{
				key:   "foo",
				value: 2,
			},
			wantErr: true,
			err:     ErrInvalidNumber,
		},
		{
			name: "key already used - error",
			cm: func() *Counter {
				cm := NewCounter()
				cm.increment("foo", 1)
				return cm
			}(),
			args: args{
				key:   "foo",
				value: 2,
			},
			wantErr: true,
			err:     ErrKeyAlreadyUsed,
		},
		{
			name: "count 1 - returns specific error",
			cm: func() *Counter {
				cm := NewCounter()
				cm.increment("foo", 1)
				return cm
			}(),
			args: args{
				key:   "foo",
				value: 1,
			},
			wantErr: true,
			err:     ErrCountRestarted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cm.increment(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				if tt.wantErr && (err != tt.err) {
					t.Errorf("CountManager.Increment() error = %v, expected %v", err, tt.err)
				}
				t.Errorf("CountManager.Increment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCountManager_Reset(t *testing.T) {
	tests := []struct {
		name string
		cm   *Counter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cm.Reset()
		})
	}
}

func TestNewCounter(t *testing.T) {
	tests := []struct {
		name string
		want *Counter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCounter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
