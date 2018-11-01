package db

import (
	"io"
	"io/ioutil"
)

func EditFile(filename string, file io.Reader) error {
	conn, err := GetDB()
	if err != nil {
		return err
	}
	defer conn.Session.Close()

	f, err := conn.GridFS("file").Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func GetFile(filename string) ([]byte, error) {
	conn, err := GetDB()
	if err != nil {
		return nil, err
	}

	f, err := conn.GridFS("file").Open(filename)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}