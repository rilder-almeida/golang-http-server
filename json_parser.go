package main

import (
	"encoding/json"
	"io/fs"
	"os"
)

func ToJsonFile(p interface{}) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = os.WriteFile(JSONFILENAME, data, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func FromJsonFile(v interface{}) error {
	data, err := os.ReadFile(JSONFILENAME)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}
