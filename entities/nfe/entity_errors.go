package nfe

import "github.com/arquivei/foundationkit/errors"

var (
	// ErrCodeDocumentNotFound is returned when a document was not found
	// in repository
	ErrCodeDocumentNotFound = errors.Code("DOCUMENT_NOT_FOUND")

	// ErrCodeGetDocument is returned when the gateway failed to get a document
	ErrCodeGetDocument = errors.Code("FAILED_GET_DOCUMENT")

	ErrCodeFailedToProcessDocument = errors.Code("FAILED_PROCESS_DOCUMENT")

	ErrCodeFailedToSaveDocument = errors.Code("FAILED_SAVE_DOCUMENT")
)
