//Package golog is a logging package that I'd like to be initially in the standard library.
//golog provides globalLog and Logger which support:
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
// - l.outFile for Trace-Info (os.Stdout by default);
// - l.errFile for Error-Fatal (os.Stderr by default).
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
var globalLog = New(customPrefixDefault, FlagsDefault)

// SetPrefix sets the output prefix for the global logger.
func SetPrefix(p string) {
	globalLog.SetPrefix(p)
}

// SetFlags sets the output flags for the global logger.
func SetFlags(f int) {
	globalLog.SetFlags(f)
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
	globalLog.SetLevel(level)
}

// SetOutput sets the output destinations for the global logger
// (different for out and err).
func SetOutput(out, err *os.File) {
	globalLog.SetOutput(out, err)
}

// Trace prints trace message to globalLog.outFile.
// Trace calls l.traceLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use trace messages for developing process to trace
// function calls.
func Trace(v ...interface{}) {
	globalLog.Trace(v...)
}

// Traceln prints trace message to globalLog.outFile.
// Traceln calls l.traceLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use trace messages for developing process to trace
// function calls.
func Traceln(v ...interface{}) {
	globalLog.Traceln(v...)
}

// Tracef prints trace message to globalLog.outFile.
// Tracef calls l.traceLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use trace messages for developing process to trace
// function calls.
func Tracef(format string, v ...interface{}) {
	globalLog.Tracef(format, v...)
}

// Debug prints debug message to globalLog.outFile.
// Debug calls l.debugLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use debug messages to debug your business logic.
func Debug(v ...interface{}) {
	globalLog.Debug(v...)
}

// Debugln prints debug message to globalLog.outFile.
// Debugln calls l.debugLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use debug messages to debug your business logic.
func Debugln(v ...interface{}) {
	globalLog.Debugln(v...)
}

// Debugf prints debug message to globalLog.outFile.
// Debugf calls l.debugLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use debug messages to debug your business logic.
func Debugf(format string, v ...interface{}) {
	globalLog.Debugf(format, v...)
}

// Info prints info message to globalLog.outFile.
// Info calls l.infoLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use info messages for common information.
func Info(v ...interface{}) {
	globalLog.Info(v...)
}

// Infoln prints info message to globalLog.outFile.
// Infoln calls l.infoLogger to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use info messages for common information.
func Infoln(v ...interface{}) {
	globalLog.Infoln(v...)
}

// Infof prints info message to globalLog.outFile.
// Infof calls l.infoLogger to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use info messages for common information.
func Infof(format string, v ...interface{}) {
	globalLog.Infof(format, v...)
}

// Print is equivalent to globalLog.Info()
func Print(v ...interface{}) {
	globalLog.Print(v...)
}

// Println is equivalent to globalLog.Infoln()
func Println(v ...interface{}) {
	globalLog.Println(v...)
}

// Printf is equivalent to globalLog.Infof()
func Printf(format string, v ...interface{}) {
	globalLog.Printf(format, v...)
}

// Warning prints warning message to globalLog.errFile.
// Warning calls l.warningLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func Warning(v ...interface{}) {
	globalLog.Warning(v...)
}

// Warningln prints warning message to globalLog.errFile.
// Warningln calls l.warningLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func Warningln(v ...interface{}) {
	globalLog.Warningln(v...)
}

// Warningf prints warning message to globalLog.errFile.
// Warningf calls l.warningLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func Warningf(format string, v ...interface{}) {
	globalLog.Warningf(format, v...)
}

// Error prints info message to globalLog.errFile.
// Error calls l.errorLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func Error(v ...interface{}) {
	globalLog.Error(v...)
}

// Errorln prints info message to globalLog.errFile.
// Errorln calls l.errorLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func Errorln(v ...interface{}) {
	globalLog.Errorln(v...)
}

// Errorf prints info message to globalLog.errFile.
// Errorf calls l.errorLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func Errorf(format string, v ...interface{}) {
	globalLog.Errorf(format, v...)
}

// Critical prints critical message to globalLog.errFile.
// Critical calls l.criticalLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use critical messages for errors which may brake
// business logic.
func Critical(v ...interface{}) {
	globalLog.Critical(v...)
}

// Criticalln prints critical message to globalLog.errFile.
// Criticalln calls l.criticalLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use critical messages for errors which may brake
// business logic.
func Criticalln(v ...interface{}) {
	globalLog.Criticalln(v...)
}

// Criticalf prints critical message to globalLog.errFile.
// Criticalf calls l.criticalLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use critical messages for errors which may brake
// business logic.
func Criticalf(format string, v ...interface{}) {
	globalLog.Criticalf(format, v...)
}

// Panic is equivalent to globalLog.Critical() followed by a call to panic().
func Panic(v ...interface{}) {
	globalLog.Panic(v...)
}

// Panicln is equivalent to globalLog.Criticalln() followed by a call to panic().
func Panicln(v ...interface{}) {
	globalLog.Panicln(v...)
}

// Panicln is equivalent to globalLog.Criticalf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	globalLog.Panicf(format, v...)
}

// Fatal prints fatal message to globalLog.errFile.
// Fatal calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatal.
func Fatal(v ...interface{}) {
	globalLog.Fatal(v...)
}

// Fatalln prints fatal message to l.errFile.
// Fatalln calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatalf.
func Fatalln(v ...interface{}) {
	globalLog.Fatalln(v...)
}

// Fatalf prints fatal message to globalLog.errFile.
// Fatalf calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatalf.
func Fatalf(format string, v ...interface{}) {
	globalLog.Fatalf(format, v...)
}

func init() {
	// necessary to provide correct call point
	globalLog.setCallDepth(4)
}
