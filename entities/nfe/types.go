package nfe

import "fmt"

type NfeDocuments []NfeDocument

var ErrNotFound = fmt.Errorf("NFe not found")

var ErrAlreadyExists = fmt.Errorf("NFe already exists")
