package golog

// Prefixes to be used for each level.
// It is possible to redefine them,
// but most probably, you don't need to do it.
var (
	PrefixTrace    = "[TRC] "
	PrefixDebug    = "[DBG] "
	PrefixInfo     = "[INF] "
	PrefixWarning  = "[WRN] "
	PrefixError    = "[ERR] "
	PrefixCritical = "[CRT] "
	PrefixPanic    = "[PNC] "
	PrefixFatal    = "[FTL] "
)
