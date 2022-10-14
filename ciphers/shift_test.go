package shift

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	type in struct {
		msg string
		key int
	}
	tests := []struct {
		name string
		in   in
		out  string
		err  error
	}{
		{
			name: "success",
			in: in{
				msg: "khan",
				key: 19,
			},
			out: "datg",
			err: nil,
		},
		{
			name: "key out of range",
			in: in{
				msg: "khan",
				key: 27,
			},
			out: "",
			err: fmt.Errorf("key must be between 0 and 26"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out, err := encrypt(tt.in.msg, tt.in.key)
			assert.Equal(t, tt.err, err)
			if diff := cmp.Diff(tt.out, out); diff != "" {
				t.Errorf("encrypt() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type in struct {
		msg string
		key int
	}
	tests := []struct {
		name string
		in   in
		out  string
		err  error
	}{
		{
			name: "success",
			in: in{
				msg: "datg",
				key: 19,
			},
			out: "khan",
			err: nil,
		},
		{
			name: "key out of range",
			in: in{
				msg: "datg",
				key: -1,
			},
			out: "",
			err: fmt.Errorf("key must be between 0 and 26"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out, err := decrypt(tt.in.msg, tt.in.key)
			assert.Equal(t, tt.err, err)
			if diff := cmp.Diff(tt.out, out); diff != "" {
				t.Errorf("decrypt() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
