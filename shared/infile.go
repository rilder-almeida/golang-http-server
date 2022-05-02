package shared

import (
	"io/fs"
	"os"
)

func FromFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			err = os.WriteFile(filename, []byte("{}"), fs.ModePerm)
			if err != nil {
				return nil, err
			}
			data, err = os.ReadFile(filename)
			if err != nil {
				return nil, err
			}
		}
	}

	return data, err
}

func ToFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
