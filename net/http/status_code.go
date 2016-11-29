package http_h

import (
	"net/http"
)

func IsSuccess(resp *http.Response) bool {
	if resp == nil {
		return false
	}

	return IsSuccessCode(resp.StatusCode)
}

func IsSuccessCode(code int) bool {
	return 200 <= code && code < 400
}
