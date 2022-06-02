package nfe

import "errors"

type NfeDocuments []NfeDocument

var ErrNotFound = errors.New("NFe not found")

var ErrAlreadyExists = errors.New("NFe already exists")
