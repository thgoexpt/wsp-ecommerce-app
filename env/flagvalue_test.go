package env

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	if GetEnv() != Testing {
		t.Errorf("expected `SOLIDENV`: %s, but get: %s", Testing, GetEnv())
	}
}