package main

import (
	"encoding/json"
	"io/fs"
	"os"
)

func ToJsonFile(filename string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func FromJsonFile(filename string, v interface{}) error {
	data, err := os.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			err = os.WriteFile(filename, []byte("{}"), fs.ModePerm)
			if err != nil {
				return err
			}
			data, err = os.ReadFile(filename)
			if err != nil {
				return err
			}
		}
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}
