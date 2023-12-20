package errorConf

import "github.com/bankonly/go-pkg/v1/stacktrace"

var (
	ErrAccessDenied1 = stacktrace.BadRequest("access_denied_1")
	ErrAccessDenied2 = stacktrace.BadRequest("access_denied_2")
	ErrAccessDenied3 = stacktrace.BadRequest("access_denied_3")
)
