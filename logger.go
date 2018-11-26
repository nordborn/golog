package golog

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// A Logger represents an active logging object that generates lines of
// output to an io.Writer. Each logging operation makes a single call to
// the Writer's Write method. A Logger can be used simultaneously from
// multiple goroutines; it guarantees to serialize access to the Writer.
// The Logger can prints messages with different levels
// If you don't know what level to use, just use Info() and Error().
// Also, it prints to different output file-like objects:
// - outFile for levels trace-info (os.Stdout by default);
// - errFile for levels warning-fatal (os.Stderr by default).
type Logger struct {
	traceLogger    *log.Logger
	debugLogger    *log.Logger
	infoLogger     *log.Logger
	warningLogger  *log.Logger
	errorLogger    *log.Logger
	criticalLogger *log.Logger
	fatalLogger    *log.Logger
	flags          int
	customPrefix   string
	level          levelType
	outFile        *os.File
	errFile        *os.File
}

func (l *Logger) updInternalLoggers() {
	l.traceLogger = log.New(l.outFile, PrefixTrace+l.customPrefix, l.flags)
	l.debugLogger = log.New(l.outFile, PrefixDebug+l.customPrefix, l.flags)
	l.infoLogger = log.New(l.outFile, PrefixInfo+l.customPrefix, l.flags)
	l.warningLogger = log.New(l.errFile, PrefixWarning+l.customPrefix, l.flags)
	l.errorLogger = log.New(l.errFile, PrefixError+l.customPrefix, l.flags)
	l.criticalLogger = log.New(l.errFile, PrefixCritical+l.customPrefix, l.flags)
	l.fatalLogger = log.New(l.errFile, PrefixFatal+l.customPrefix, l.flags)
}

func (l *Logger) updOutputsToLevel() {
	if l.level == LevelDebug {
		l.traceLogger.SetOutput(ioutil.Discard)
	}
	if l.level == LevelInfo {
		l.traceLogger.SetOutput(ioutil.Discard)
		l.debugLogger.SetOutput(ioutil.Discard)
	}
	if l.level == LevelError {
		l.traceLogger.SetOutput(ioutil.Discard)
		l.debugLogger.SetOutput(ioutil.Discard)
		l.infoLogger.SetOutput(ioutil.Discard)
	}
}

// New creates new logger.
// Use flags==-1 to set default flags
func New(customPrefix string, flags int) *Logger {
	if flags < 0 || flags > 255 {
		flags = FlagsDefault
	}
	l := Logger{}
	l.outFile = OutDefault
	l.errFile = ErrDefault
	l.SetPrefix(customPrefix)
	l.SetFlags(flags)
	return &l
}

// SetLevel sets the output level for the logger.
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
func (l *Logger) SetLevel(level levelType) {
	l.level = level
	l.updInternalLoggers()
	l.updOutputsToLevel()
}

// SetPrefix sets the output prefix for the logger.
func (l *Logger) SetPrefix(p string) {
	p = strings.Trim(p, " ")
	if p != "" {
		p += " "
	}
	l.customPrefix = p
	l.updInternalLoggers()
	l.updOutputsToLevel()
}

// SetFlags sets the output flags for the logger.
func (l *Logger) SetFlags(f int) {
	if f < 0 || f > 255 {
		f = FlagsDefault
	}
	l.flags = f
	l.updInternalLoggers()
	l.updOutputsToLevel()
}

// SetOutput sets the output destinations for the logger
// (different for out and err).
func (l *Logger) SetOutput(out, err *os.File) {
	l.outFile = out
	l.errFile = err
	l.updInternalLoggers()
	l.updOutputsToLevel()
}

// Trace prints trace message to l.outFile.
// Trace calls l.traceLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use trace messages for developing process to trace
// function calls.
func (l *Logger) Trace(v ...interface{}) {
	l.traceLogger.Print(v...)
}

// Traceln prints trace message to l.outFile.
// Traceln calls l.traceLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use trace messages for developing process to trace
// function calls.
func (l *Logger) Traceln(v ...interface{}) {
	l.traceLogger.Println(v...)
}

// Tracef prints trace message to l.outFile.
// Tracef calls l.traceLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use trace messages for developing process to trace
// function calls.
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.traceLogger.Printf(format, v...)
}

// Debug prints debug message to l.outFile.
// Debug calls l.debugLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use debug messages to debug your business logic.
func (l *Logger) Debug(v ...interface{}) {
	l.debugLogger.Print(v...)
}

// Debugln prints debug message to l.outFile.
// Debugln calls l.debugLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use debug messages to debug your business logic.
func (l *Logger) Debugln(v ...interface{}) {
	l.debugLogger.Println(v...)
}

// Debugf prints debug message to l.outFile.
// Debugf calls l.debugLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use debug messages to debug your business logic.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debugLogger.Printf(format, v...)
}

// Info prints info message to l.outFile.
// Info calls l.infoLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use info messages for common information.
func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Print(v...)
}

// Infoln prints info message to l.outFile.
// Infoln calls l.infoLogger to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use info messages for common information.
func (l *Logger) Infoln(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Infof prints info message to l.outFile.
// Infof calls l.infoLogger to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use info messages for common information.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Print is equivalent to l.Info()
func (l *Logger) Print(v ...interface{}) {
	l.infoLogger.Print(v...)
}

// Println is equivalent to l.Infoln()
func (l *Logger) Println(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Printf is equivalent to l.Infof()
func (l *Logger) Printf(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Warning prints warning message to l.errFile.
// Warning calls l.warningLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func (l *Logger) Warning(v ...interface{}) {
	l.warningLogger.Print(v...)
}

// Warningln prints warning message to l.errFile.
// Warningln calls l.warningLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func (l *Logger) Warningln(v ...interface{}) {
	l.warningLogger.Println(v...)
}

// Warningf prints warning message to l.errFile.
// Warningf calls l.warningLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warningLogger.Printf(format, v...)
}

// Error prints info message to l.errFile.
// Error calls l.errorLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Print(v...)
}

// Errorln prints info message to l.errFile.
// Errorln calls l.errorLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func (l *Logger) Errorln(v ...interface{}) {
	l.errorLogger.Println(v...)
}

// Errorf prints info message to l.errFile.
// Errorf calls l.errorLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Critical prints critical message to l.errFile.
// Critical calls l.criticalLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use critical messages for errors which may brake
// business logic.
func (l *Logger) Critical(v ...interface{}) {
	l.criticalLogger.Print(v...)
}

// Criticalln prints critical message to l.errFile.
// Criticalln calls l.criticalLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use critical messages for errors which may brake
// business logic.
func (l *Logger) Criticalln(v ...interface{}) {
	l.criticalLogger.Println(v...)
}

// Criticalf prints critical message to l.errFile.
// Criticalf calls l.criticalLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use critical messages for errors which may brake
// business logic.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	l.criticalLogger.Printf(format, v...)
}

// Panic is equivalent to l.Critical() followed by a call to panic().
func (l *Logger) Panic(v ...interface{}) {
	l.criticalLogger.Panic(v...)
}

// Panicln is equivalent to l.Criticalln() followed by a call to panic().
func (l *Logger) Panicln(v ...interface{}) {
	l.criticalLogger.Panicln(v...)
}

// Panicln is equivalent to l.Criticalf() followed by a call to panic().
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.criticalLogger.Panicf(format, v...)
}

// Fatal prints fatal message to l.errFile.
// Fatal calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatal.
func (l *Logger) Fatal(v ...interface{}) {
	l.fatalLogger.Fatal(v...)
}

// Fatalln prints fatal message to l.errFile.
// Fatalln calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatalln.
func (l *Logger) Fatalln(v ...interface{}) {
	l.fatalLogger.Fatalln(v...)
}

// Fatalf prints fatal message to l.errFile.
// Fatalf calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatalf.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.fatalLogger.Fatalf(format, v...)
}
