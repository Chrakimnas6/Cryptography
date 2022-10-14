package rsa

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
		err  error
	}{
		{
			name: "success",
			in: in{
				msg: "khan",
				p:   53,
				q:   59,
			},
			out: "khan",
			err: nil,
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
