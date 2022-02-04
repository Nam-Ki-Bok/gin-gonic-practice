package security

import (
	"net/http"
	"strings"
)

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	// Bearer Token..
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
