package nfe

import fkerrors "github.com/arquivei/foundationkit/errors"

var (
	// ErrCodeDocumentNotFound is returned when a document was not found
	// in repository
	ErrCodeDocumentNotFound = fkerrors.Code("DOCUMENT_NOT_FOUND")

	// ErrCodeGetDocument is returned when the gateway failed to get a document

	ErrCodeProcessDocument = fkerrors.Code("FAILED_PROCESS_DOCUMENT")

	ErrCodeSaveDocument = fkerrors.Code("FAILED_SAVE_DOCUMENT")

	ErrCodeGetDocument = fkerrors.Code("FAILED_GET_DOCUMENT")
)
