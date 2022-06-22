package nfe

import "errors"

type NFeDocuments []NFeDocument

var ErrNotFound = errors.New("NFe not found")

var ErrAlreadyExists = errors.New("NFe already exists")
