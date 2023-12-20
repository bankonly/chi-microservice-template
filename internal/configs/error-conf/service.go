package errorConf

import "github.com/bankonly/go-pkg/v1/stacktrace"

var (
	ErrBadParemeter = stacktrace.BadRequest("bad_request")
)
