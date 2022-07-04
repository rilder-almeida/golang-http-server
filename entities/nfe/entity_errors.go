package nfe

import (
	"errors"

	fkerrors "github.com/arquivei/foundationkit/errors"
)

var (
	// ErrCodeDocumentNotFound is returned when a document was not found
	// in repository
	ErrCodeDocumentNotFound = fkerrors.Code("DOCUMENT_NOT_FOUND")

	// ErrCodeGetDocument is returned when the gateway failed to get a document

	ErrCodeProcessDocument = fkerrors.Code("FAILED_PROCESS_DOCUMENT")

	ErrCodeSaveDocument = fkerrors.Code("FAILED_SAVE_DOCUMENT")

	ErrCodeGetDocument = fkerrors.Code("FAILED_GET_DOCUMENT")
)

var (
	ErrDocumentNotFound = errors.New("DOCUMENT_NOT_FOUND")

	ErrProcessDocument = errors.New("FAILED_PROCESS_DOCUMENT")

	ErrSaveDocument = errors.New("FAILED_SAVE_DOCUMENT")

	ErrEmptyXML = errors.New("EMPTY_XML")

	ErrEmptyId = errors.New("EMPTY_ID")
)
