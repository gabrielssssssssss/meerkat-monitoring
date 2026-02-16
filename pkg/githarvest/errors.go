package githarvest

import "errors"

var (
	ErrCreateRequest = errors.New("failed to create http request")

	ErrExecRequest = errors.New("failed to execute http request")

	ErrReadBody = errors.New("failed to read response body")

	ErrDecodeJSON = errors.New("failed to decode json response")
)
