package env

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	if GetEnv() != Testing && GetEnv() != CI {
		t.Errorf("expected `SOLIDENV`: %s or %s, but get: %s", Testing, CI, GetEnv())
	}
}

func TestGetPort(t *testing.T) {
	if GetPort() != os.Getenv("PORT") {
		t.Errorf("expected `PORT`: %s, but get: %s", os.Getenv("PORT"), GetPort())
	}
}