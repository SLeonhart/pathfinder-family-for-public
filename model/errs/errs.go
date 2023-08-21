package errs

import "errors"

var (
	NetworkError    = errors.New("NetworkError")
	InvalidResponse = errors.New("InvalidResponse")
	NonOkStatusCode = errors.New("NonOK status code returned")
	InvalidRequest  = errors.New("InvalidRequest")
	ProxyError      = errors.New("ProxyError")
)
