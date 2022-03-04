package main

import (
	"encoding/json"
	"os"
)

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func ReadJsonFromFile(filename string, v interface{}) error {
	if FileExists(filename) {
		err := CreateFile(filename)
		if err != nil {
			return err
		}
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []byte
	file.Read(data)
	return Unmarshal(data, v)
}

func WriteJsonToFile(filename string, v interface{}) error {
	if FileExists(filename) {
		err := CreateFile(filename)
		if err != nil {
			return err
		}
	}
	err := TruncateFile(filename)
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := Marshal(v)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}

func CreateFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func TruncateFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	return nil
}
