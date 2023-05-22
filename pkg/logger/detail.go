package logger

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

type Detail struct {
	Request      string         `json:"request,omitempty"`
	Address      string         `json:"address,omitempty"`
	Backtrace    string         `json:"backtrace,omitempty"`
	Response     string         `json:"response,omitempty"`
	ResponseCode int            `json:"response_code,omitempty"`
	Fields       map[string]any `json:"fields,omitempty"`
}

func GetBacktrace(err error) string {
	stack := string(debug.Stack())

	if err != nil {
		stack = fmt.Sprintf("%s\n%s", err.Error(), stack)
	}

	return stack
}

func GetFieldsForServer(r *http.Request) map[string]interface{} {
	return map[string]interface{}{
		"method":     r.Method,
		"query":      r.URL.RawQuery,
		"ip":         r.RemoteAddr,
		"user-agent": r.UserAgent(),
	}
}
