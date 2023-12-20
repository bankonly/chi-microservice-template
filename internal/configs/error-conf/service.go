package errorConf

import "github.com/bankonly/go-pkg/v1/stacktrace"

var (
	ErrBadParameter = stacktrace.BadRequest("bad_request")
)
