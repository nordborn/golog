package golog

// Prefixes to be used for each level.
// It is possible to redefine them,
// but most probably, you don't need to do it.
var (
	PrefixTrace    = "TRACE: "
	PrefixDebug    = "DEBUG: "
	PrefixInfo     = "INFO: "
	PrefixWarning  = "WARNING: "
	PrefixError    = "ERROR: "
	PrefixCritical = "CRITICAL: "
	PrefixFatal    = "FATAL: "
)
