package transparency

import "errors"

var (
	ErrUnknowEntryType = errors.New("unknow entry type")
	ErrFailedFetch     = errors.New("failed to fetch url")
	ErrFailedRead      = errors.New("failed read io body")
	ErrJsonMarshal     = errors.New("failed to convert bytes to json")
)
