//Package golog is a logging package that I'd like to be initially in the standard library.
//golog provides loggerGlobal and Logger which support:
//1. Levels:
//- trace;
//- debug;
//- info;
//- warning;
//- error;
//- critical.
//2. Different outputs:
//- for info-like messages (Trace, Debug, Info), they use os.Stdout by default;
//- for error-like messages (Warning, Error, Critical, Panic, Fatal) they use os.Stderr by default.
//
//You can set:
//1. logging level (LevelTrace, LevelDebug, LevelInfo, LevelWarning, LevelError, LevelCritical);
//2. custom prefix (e.g. "[myapp]: ") additionally to level prefixes ("[main]: " by default);
//3. output file-like interfaces:
// - l.outWriter for Trace-Info (os.Stdout by default);
// - l.errWriter for Error-Fatal (os.Stderr by default).
//
//You can set flag similar to "log" from standard library for time and file information
//(default provides "2018/11/26 16:57:49 golog.go:61").
//
//Also, you can change level prefixes (defaults are "TRACE: ", "DEBUG: " etc.) but don't do it
//if you don't need it really.
//
//Tip: in common usage if you don't know which messages you should use, use Infoln() and Errorln().
//
//So, common message using `golog.Infoln("Started")` will be:
//INFO: [main]: 2018/11/26 16:57:49 main.go:61: Started

package golog

import "os"

// will be used in package-level logging functions
var loggerGlobal = New(customPrefixDefault, FlagsDefault)

// SetPrefix sets the output prefix for the global logger.
func SetPrefix(p string) {
	loggerGlobal.SetPrefix(p)
}

// SetFlags sets the output flags for the global logger.
func SetFlags(f int) {
	loggerGlobal.SetFlags(f)
}

// SetLevel sets the output level for the global logger.
// It allows to suppress messages from unimportant levels.
// Levels hierarchy:
// - trace;
// - debug;
// - info;
// - warning;
// - error;
// - critical.
// Use:
// LevelTrace - to display all messages;
// LevelDebug - to display debug messages and above;
// LevelInfo - to display info messages and above;
// LevelWarning - to display warning messages and above;
// LevelError - to display error messages and above.
// Default level: LevelTrace
func SetLevel(level levelType) {
	loggerGlobal.SetLevel(level)
}

// SetOutput sets the output destinations for the global logger
// (different for out and err).
func SetOutput(out, err *os.File) {
	loggerGlobal.SetOutput(out, err)
}

// Trace prints trace message to loggerGlobal.outWriter.
// Trace calls l.traceLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use trace messages for developing process to trace
// function calls.
func Trace(v ...interface{}) {
	loggerGlobal.Trace(v...)
}

// Traceln prints trace message to loggerGlobal.outWriter.
// Traceln calls l.traceLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use trace messages for developing process to trace
// function calls.
func Traceln(v ...interface{}) {
	loggerGlobal.Traceln(v...)
}

// Tracef prints trace message to loggerGlobal.outWriter.
// Tracef calls l.traceLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use trace messages for developing process to trace
// function calls.
func Tracef(format string, v ...interface{}) {
	loggerGlobal.Tracef(format, v...)
}

// Debug prints debug message to loggerGlobal.outWriter.
// Debug calls l.debugLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use debug messages to debug your business logic.
func Debug(v ...interface{}) {
	loggerGlobal.Debug(v...)
}

// Debugln prints debug message to loggerGlobal.outWriter.
// Debugln calls l.debugLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use debug messages to debug your business logic.
func Debugln(v ...interface{}) {
	loggerGlobal.Debugln(v...)
}

// Debugf prints debug message to loggerGlobal.outWriter.
// Debugf calls l.debugLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use debug messages to debug your business logic.
func Debugf(format string, v ...interface{}) {
	loggerGlobal.Debugf(format, v...)
}

// Info prints info message to loggerGlobal.outWriter.
// Info calls l.infoLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use info messages for common information.
func Info(v ...interface{}) {
	loggerGlobal.Info(v...)
}

// Infoln prints info message to loggerGlobal.outWriter.
// Infoln calls l.infoLogger to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use info messages for common information.
func Infoln(v ...interface{}) {
	loggerGlobal.Infoln(v...)
}

// Infof prints info message to loggerGlobal.outWriter.
// Infof calls l.infoLogger to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use info messages for common information.
func Infof(format string, v ...interface{}) {
	loggerGlobal.Infof(format, v...)
}

// Print is equivalent to loggerGlobal.Info()
func Print(v ...interface{}) {
	loggerGlobal.Print(v...)
}

// Println is equivalent to loggerGlobal.Infoln()
func Println(v ...interface{}) {
	loggerGlobal.Println(v...)
}

// Printf is equivalent to loggerGlobal.Infof()
func Printf(format string, v ...interface{}) {
	loggerGlobal.Printf(format, v...)
}

// Warning prints warning message to loggerGlobal.errWriter.
// Warning calls l.warningLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func Warning(v ...interface{}) {
	loggerGlobal.Warning(v...)
}

// Warningln prints warning message to loggerGlobal.errWriter.
// Warningln calls l.warningLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func Warningln(v ...interface{}) {
	loggerGlobal.Warningln(v...)
}

// Warningf prints warning message to loggerGlobal.errWriter.
// Warningf calls l.warningLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func Warningf(format string, v ...interface{}) {
	loggerGlobal.Warningf(format, v...)
}

// Error prints info message to loggerGlobal.errWriter.
// Error calls l.errorLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func Error(v ...interface{}) {
	loggerGlobal.Error(v...)
}

// Errorln prints info message to loggerGlobal.errWriter.
// Errorln calls l.errorLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func Errorln(v ...interface{}) {
	loggerGlobal.Errorln(v...)
}

// Errorf prints info message to loggerGlobal.errWriter.
// Errorf calls l.errorLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func Errorf(format string, v ...interface{}) {
	loggerGlobal.Errorf(format, v...)
}

// Critical prints critical message to loggerGlobal.errWriter.
// Critical calls l.criticalLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use critical messages for errors which may brake
// business logic.
func Critical(v ...interface{}) {
	loggerGlobal.Critical(v...)
}

// Criticalln prints critical message to loggerGlobal.errWriter.
// Criticalln calls l.criticalLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use critical messages for errors which may brake
// business logic.
func Criticalln(v ...interface{}) {
	loggerGlobal.Criticalln(v...)
}

// Criticalf prints critical message to loggerGlobal.errWriter.
// Criticalf calls l.criticalLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use critical messages for errors which may brake
// business logic.
func Criticalf(format string, v ...interface{}) {
	loggerGlobal.Criticalf(format, v...)
}

// Panic is equivalent to loggerGlobal.Critical() followed by a call to panic().
func Panic(v ...interface{}) {
	loggerGlobal.Panic(v...)
}

// Panicln is equivalent to loggerGlobal.Criticalln() followed by a call to panic().
func Panicln(v ...interface{}) {
	loggerGlobal.Panicln(v...)
}

// Panicln is equivalent to loggerGlobal.Criticalf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	loggerGlobal.Panicf(format, v...)
}

// Fatal prints fatal message to loggerGlobal.errWriter.
// Fatal calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatal.
func Fatal(v ...interface{}) {
	loggerGlobal.Fatal(v...)
}

// Fatalln prints fatal message to l.errWriter.
// Fatalln calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatalf.
func Fatalln(v ...interface{}) {
	loggerGlobal.Fatalln(v...)
}

// Fatalf prints fatal message to loggerGlobal.errWriter.
// Fatalf calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatalf.
func Fatalf(format string, v ...interface{}) {
	loggerGlobal.Fatalf(format, v...)
}

func init() {
	// necessary to provide correct call point
	loggerGlobal.setCallDepth(4)
}
