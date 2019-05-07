package utils

import (
	"net/http"
)

func Abort(w http.ResponseWriter, message string, code int) {
	http.Error(w, message, code)
}
