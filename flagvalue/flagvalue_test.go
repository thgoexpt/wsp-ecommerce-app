package flagvalue

import (
	"flag"
	"testing"
)

func TestGetEnv(t *testing.T) {
	var env string

	defaultVal := "TESTING"
	flag.StringVar(&env, "env", defaultVal, "Indicates the program runs on the TESTING or PRODUCTION environment")
	flag.Parse()

	if GetEnv() != defaultVal {
		t.Errorf("expected: %s, but get: %s", GetEnv(), defaultVal)
	}

	expected := "PRODUCTION"
	flag.Set("env", expected)
	if GetEnv() != expected {
		t.Errorf("expected: %s, but get: %s", GetEnv(), expected)
	}
}

func TestGetDBHost(t *testing.T) {
	var dbhost string

	defaultVal := "localhost"
	flag.StringVar(&dbhost, "dbhost", defaultVal, "Database host")
	flag.Parse()

	if GetDBHost() != defaultVal {
		t.Errorf("expected: %s, but get: %s", GetDBHost(), defaultVal)
	}

	expected := "192.168.1.2"
	flag.Set("dbhost", expected)
	if GetDBHost() != expected {
		t.Errorf("expected: %s, but get: %s", GetDBHost(), expected)
	}
}

func TestGetDBPort(t *testing.T) {
	var dbport string

	defaultVal := "27017"
	flag.StringVar(&dbport, "dbport", defaultVal, "Database port")
	flag.Parse()

	if GetDBPort() != defaultVal {
		t.Errorf("expected: %s, but get: %s", GetDBPort(), defaultVal)
	}

	expected := "9999"
	flag.Set("dbport", expected)
	if GetDBPort() != expected {
		t.Errorf("expected: %s, but get: %s", GetDBPort(), expected)
	}
}

func TestGetDBName(t *testing.T) {
	var dbname string

	defaultVal := "mongotest"
	flag.StringVar(&dbname, "dbname", defaultVal, "Database name")
	flag.Parse()

	if GetDBName() != defaultVal {
		t.Errorf("expected: %s, but get: %s", GetDBName(), defaultVal)
	}

	expected := "mongo"
	flag.Set("dbname", expected)
	if GetDBName() != expected {
		t.Errorf("expected: %s, but get: %s", GetDBName(), expected)
	}
}
