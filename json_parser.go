package main

import (
	"encoding/json"
	"io/fs"
	"os"
)

func Marshal(v *interface{}) ([]byte, error) {
	return json.Marshal(&v)
}

func Unmarshal(data []byte, v *interface{}) error {
	return json.Unmarshal(data, &v)
}

func ReadJsonFromFile(filename string, v interface{}) error {
	if FileExists(filename) {
		err := CreateFile(filename)
		if err != nil {
			return err
		}
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return Unmarshal(data, &v)
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

	data, err := Marshal(&v)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, fs.ModePerm)
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
	err := os.WriteFile(filename, []byte{}, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
