package errhandler

import (
	"net/http"

	"github.com/rhythin/bookspot/services/shared/customlogger"
)

type httpHandler func(w http.ResponseWriter, r *http.Request) (err error)

func HttpErrorHandler(handler func(w http.ResponseWriter, r *http.Request) (err error)) http.HandlerFunc {
	return httpHandler(handler).ServeHTTP
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := customlogger.S()

	// call the handler function
	err := h(w, r)
	// if no errors return
	if err == nil {
		return
	}

	// if error is not of type CustomError, set it to unknown error
	httpErr, ok := err.(*CustomError)
	if !ok {
		logger.Warnw("Unhandled error", "error", err)
		httpErr = unknownError.SetError(err)
	}

	// get error response body
	body, err := httpErr.ErrorResponseBody()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write error response code and body
	w.WriteHeader(httpErr.statusCode)
	w.Write(body) //nolint
}
