package golog

import (
	"fmt"
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
	panicLogger    *log.Logger
	fatalLogger    *log.Logger
	flags          int
	customPrefix   string
	level          levelType
	outFile        *os.File
	errFile        *os.File
	calldepth      int
}

func (l *Logger) updInternalLoggers() {
	l.traceLogger = log.New(l.outFile, PrefixTrace+l.customPrefix, l.flags)
	l.debugLogger = log.New(l.outFile, PrefixDebug+l.customPrefix, l.flags)
	l.infoLogger = log.New(l.outFile, PrefixInfo+l.customPrefix, l.flags)
	l.warningLogger = log.New(l.errFile, PrefixWarning+l.customPrefix, l.flags)
	l.errorLogger = log.New(l.errFile, PrefixError+l.customPrefix, l.flags)
	l.criticalLogger = log.New(l.errFile, PrefixCritical+l.customPrefix, l.flags)
	l.panicLogger = log.New(l.errFile, PrefixPanic+l.customPrefix, l.flags)
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
	if l.level == LevelWarning {
		l.traceLogger.SetOutput(ioutil.Discard)
		l.debugLogger.SetOutput(ioutil.Discard)
		l.infoLogger.SetOutput(ioutil.Discard)
	}
	if l.level == LevelError {
		l.traceLogger.SetOutput(ioutil.Discard)
		l.debugLogger.SetOutput(ioutil.Discard)
		l.infoLogger.SetOutput(ioutil.Discard)
		l.warningLogger.SetOutput(ioutil.Discard)
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
	l.calldepth = 3 // as for log.Logger
	l.SetPrefix(customPrefix)
	l.SetFlags(flags)
	return &l
}

// only for global logger
func (l *Logger) setCallDepth(v int) {
	if v < 1 {
		v = 1
	}
	l.calldepth = v
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

func outputln(calldepth int, ll *log.Logger, v ...interface{}) {
	ll.Output(calldepth, fmt.Sprintln(v...))
}

func output(calldepth int, ll *log.Logger, v ...interface{}) {
	ll.Output(calldepth, fmt.Sprint(v...))
}

func outputf(calldepth int, ll *log.Logger, format string, v ...interface{}) {
	ll.Output(calldepth, fmt.Sprintf(format, v...))
}

// Trace prints trace message to l.outFile.
// Trace calls l.traceLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use trace messages for developing process to trace
// function calls.
func (l *Logger) Trace(v ...interface{}) {
	output(l.calldepth, l.traceLogger, v...)
}

// Traceln prints trace message to l.outFile.
// Traceln calls l.traceLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use trace messages for developing process to trace
// function calls.
func (l *Logger) Traceln(v ...interface{}) {
	outputln(l.calldepth, l.traceLogger, v...)
}

// Tracef prints trace message to l.outFile.
// Tracef calls l.traceLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use trace messages for developing process to trace
// function calls.
func (l *Logger) Tracef(format string, v ...interface{}) {
	outputf(l.calldepth, l.traceLogger, format, v...)
}

// Debug prints debug message to l.outFile.
// Debug calls l.debugLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use debug messages to debug your business logic.
func (l *Logger) Debug(v ...interface{}) {
	output(l.calldepth, l.debugLogger, v...)
}

// Debugln prints debug message to l.outFile.
// Debugln calls l.debugLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use debug messages to debug your business logic.
func (l *Logger) Debugln(v ...interface{}) {
	outputln(l.calldepth, l.debugLogger, v...)
}

// Debugf prints debug message to l.outFile.
// Debugf calls l.debugLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use debug messages to debug your business logic.
func (l *Logger) Debugf(format string, v ...interface{}) {
	outputf(l.calldepth, l.debugLogger, format, v...)
}

// Info prints info message to l.outFile.
// Info calls l.infoLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use info messages for common information.
func (l *Logger) Info(v ...interface{}) {
	output(l.calldepth, l.infoLogger, v...)
}

// Infoln prints info message to l.outFile.
// Infoln calls l.infoLogger to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use info messages for common information.
func (l *Logger) Infoln(v ...interface{}) {
	outputln(l.calldepth, l.infoLogger, v...)
}

// Infof prints info message to l.outFile.
// Infof calls l.infoLogger to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use info messages for common information.
func (l *Logger) Infof(format string, v ...interface{}) {
	outputf(l.calldepth, l.infoLogger, format, v...)
}

// Print is equivalent to l.Info()
func (l *Logger) Print(v ...interface{}) {
	output(l.calldepth, l.infoLogger, v...)
}

// Println is equivalent to l.Infoln()
func (l *Logger) Println(v ...interface{}) {
	outputln(l.calldepth, l.infoLogger, v...)
}

// Printf is equivalent to l.Infof()
func (l *Logger) Printf(format string, v ...interface{}) {
	outputf(l.calldepth, l.infoLogger, format, v...)
}

// Warning prints warning message to l.errFile.
// Warning calls l.warningLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func (l *Logger) Warning(v ...interface{}) {
	output(l.calldepth, l.warningLogger, v...)
}

// Warningln prints warning message to l.errFile.
// Warningln calls l.warningLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func (l *Logger) Warningln(v ...interface{}) {
	outputln(l.calldepth, l.warningLogger, v...)
}

// Warningf prints warning message to l.errFile.
// Warningf calls l.warningLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use warning messages for handled errors which don't brake
// business logic but should be noted (mostly for developers).
func (l *Logger) Warningf(format string, v ...interface{}) {
	outputf(l.calldepth, l.warningLogger, format, v...)
}

// Error prints info message to l.errFile.
// Error calls l.errorLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func (l *Logger) Error(v ...interface{}) {
	output(l.calldepth, l.errorLogger, v...)
}

// Errorln prints info message to l.errFile.
// Errorln calls l.errorLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func (l *Logger) Errorln(v ...interface{}) {
	outputln(l.calldepth, l.errorLogger, v...)
}

// Errorf prints info message to l.errFile.
// Errorf calls l.errorLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use error messages for errors which mostly don't brake
// business logic.
func (l *Logger) Errorf(format string, v ...interface{}) {
	outputf(l.calldepth, l.errorLogger, format, v...)
}

// Critical prints critical message to l.errFile.
// Critical calls l.criticalLogger.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// Tip: use critical messages for errors which may brake
// business logic.
func (l *Logger) Critical(v ...interface{}) {
	output(l.calldepth, l.criticalLogger, v...)
}

// Criticalln prints critical message to l.errFile.
// Criticalln calls l.criticalLogger.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// Tip: use critical messages for errors which may brake
// business logic.
func (l *Logger) Criticalln(v ...interface{}) {
	outputln(l.calldepth, l.criticalLogger, v...)
}

// Criticalf prints critical message to l.errFile.
// Criticalf calls l.criticalLogger.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// Tip: use critical messages for errors which may brake
// business logic.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	outputf(l.calldepth, l.criticalLogger, format, v...)
}

// Panic is equivalent to l.Critical() followed by a call to panic().
func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	output(l.calldepth, l.panicLogger, s)
	panic(s)
}

// Panicln is equivalent to l.Criticalln() followed by a call to panic().
func (l *Logger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	outputln(l.calldepth, l.panicLogger, s)
	panic(s)
}

// Panicln is equivalent to l.Criticalf() followed by a call to panic().
func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	outputf(l.calldepth, l.panicLogger, format, s)
	panic(s)
}

// Fatal prints fatal message to l.errFile.
// Fatal calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatal.
func (l *Logger) Fatal(v ...interface{}) {
	output(l.calldepth, l.fatalLogger, v...)
	os.Exit(1)
}

// Fatalln prints fatal message to l.errFile.
// Fatalln calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatalln.
func (l *Logger) Fatalln(v ...interface{}) {
	outputln(l.calldepth, l.fatalLogger, v...)
	os.Exit(1)
}

// Fatalf prints fatal message to l.errFile.
// Fatalf calls l.fatalLogger.Print to print to the logger
// followed by a call to os.Exit(1).
// Note: recover() can't intercept Fatalf.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	outputf(l.calldepth, l.fatalLogger, format, v...)
	os.Exit(1)
}
