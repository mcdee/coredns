package ens

import (
	"testing"

	"github.com/mholt/caddy"
)

func TestEnsParse(t *testing.T) {
	tests := []struct {
		inputFileRules     string
		shouldErr          bool
		expectedConnection string
	}{
		{
			`ens {
			}`,
			true,
			"",
		},
		{
			`ens {
				connection
			}`,
			true,
			"",
		},
		{
			`ens {
				connection /home/test/.ethereum/geth.ipc
			}`,
			false,
			"/home/test/.ethereum/geth.ipc",
		},
		{
			`ens {
					connection http://localhost:8545/
			}`,
			false,
			"http://localhost:8545/",
		},
	}

	for i, test := range tests {
		c := caddy.NewTestController("dns", test.inputFileRules)
		actualConnection, err := ensParse(c)

		if err == nil && test.shouldErr {
			t.Fatalf("Test %d expected errors, but got no error", i)
		} else if err != nil && !test.shouldErr {
			t.Fatalf("Test %d expected no errors, but got '%v'", i, err)
		} else {
			if actualConnection != test.expectedConnection {
				t.Fatalf("Test %d expected %v, got %v", i, test.expectedConnection, actualConnection)
			}
		}
	}
}
