package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveFlag(t *testing.T) {
	tests := []struct {
		desc       string
		envName    string
		envVal     string
		setEnv     bool
		defaultVal string
		testFunc   func(t *testing.T, got *string)
	}{
		{
			desc:       "happy path, use default val",
			envName:    "DB_HOST",
			envVal:     "some_value",
			setEnv:     false,
			defaultVal: "localhost",
			testFunc: func(t *testing.T, got *string) {
				assert.Equal(t, *got, "localhost")
			},
		},
		{
			desc:       "happy path, use env val",
			envName:    "DB_HOST",
			envVal:     "some_value",
			setEnv:     true,
			defaultVal: "localhost",
			testFunc: func(t *testing.T, got *string) {
				assert.Equal(t, *got, "some_value")
			},
		},
	}
	for _, tc := range tests {
		if tc.setEnv {
			os.Setenv(tc.envName, tc.envVal)
		}
		got := resolveFlag(tc.envName, tc.defaultVal)
		tc.testFunc(t, got)
	}
}
