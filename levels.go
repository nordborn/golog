package golog

type levelType int

// Levels hierarchy:
// - trace;
// - debug;
// - info;
// - warning;
// - error;
// - critical.
// Use in SetLevel():
// LevelTrace - to display all messages;
// LevelDebug - to display debug messages and above;
// LevelInfo - to display info messages and above;
// LevelWarning - to display warning messages and above;
// LevelError - to display error messages and above.
// Default level: LevelTrace
const (
	LevelTrace levelType = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)
