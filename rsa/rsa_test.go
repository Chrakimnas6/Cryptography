package rsa

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestRSA(t *testing.T) {
	type in struct {
		msg  string
		p, q int
	}
	tests := []struct {
		name string
		in   in
		out  string
	}{
		{
			name: "success",
			in: in{
				msg: "khan",
				p:   53,
				q:   59,
			},
			out: "khan",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			pub, priv := generateKeys(tt.in.p, tt.in.q)
			out := decrypt(encrypt(tt.in.msg, pub), priv)
			if diff := cmp.Diff(tt.out, out); diff != "" {
				t.Errorf("decrypt(encrypt()) mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRSAUsingLibrary(t *testing.T) {
	tests := []struct {
		name string
		msg  string
	}{
		{
			name: "success",
			msg:  "secret mmessage",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			pub, priv, err := generateKeysUsingLibrary()
			assert.NoError(t, err)
			msg, err := encryptUsingLibrary(tt.msg, pub)
			assert.NoError(t, err)
			out, err := decryptUsingLibrary(msg, priv)
			assert.NoError(t, err)
			if diff := cmp.Diff(tt.msg, out); diff != "" {
				t.Errorf("decryptUsingLibrary() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
