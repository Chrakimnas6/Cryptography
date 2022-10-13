package shift

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
	}{
		{
			name: "success",
			in: in{
				msg: "hello",
				key: 27,
			},
			out: "ifmmp",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := Encrypt(tt.in.msg, tt.in.key)
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
	}{
		{
			name: "success",
			in: in{
				msg: "ifmmp",
				key: 27,
			},
			out: "hello",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := Decrypt(tt.in.msg, tt.in.key)
			if diff := cmp.Diff(tt.out, out); diff != "" {
				t.Errorf("decrypt() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
