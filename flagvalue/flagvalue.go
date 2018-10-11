package flagvalue

import "flag"

func GetEnv() string {
	flag.Parse()
	return flag.Lookup("env").Value.(flag.Getter).Get().(string)
}

func GetDBHost() string {
	flag.Parse()
	return flag.Lookup("dbhost").Value.(flag.Getter).Get().(string)
}

func GetDBPort() string {
	flag.Parse()
	return flag.Lookup("dbport").Value.(flag.Getter).Get().(string)
}

func GetDBName() string {
	flag.Parse()
	return flag.Lookup("dbname").Value.(flag.Getter).Get().(string)
}
