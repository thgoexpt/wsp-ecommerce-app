package env

import "os"

var Production = "PRODUCTION"
var Testing = "TESTING"
var CI = "CI"

func GetEnv() string {
	env := os.Getenv("SOLID_ENV")
	if env != Production && env != CI {
		env = Testing
	}
	return env
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetMongoURI() string {
	return os.Getenv("MONGODB_URI")
}