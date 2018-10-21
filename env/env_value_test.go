package env

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	if GetEnv() != Testing && GetEnv() != CI {
		t.Errorf("expected `SOLIDENV`: %s or %s, but get: %s", Testing, CI, GetEnv())
	}
}