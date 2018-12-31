package golog

import (
	"log"
	"testing"
)

func TestStdLog(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			Criticalln("Recovered:", r)
		}
	}()
	SetLevel(LevelInfo)
	Infoln("Initial prefix")
	SetPrefix("mylog:")

	Traceln("You shouldn't see it")
	Infoln("Custom prefix")
	Warningln("Warning")
	Errorln("Error")
	Panicln("Panic")

	SetFlags(log.Ltime | log.Lshortfile)
	Errorln("Short time")
	SetLevel(LevelTrace)
	Traceln("Trace")
}

func TestLogger(t *testing.T) {
	l := New("customlog:", -1)
	l.SetLevel(LevelInfo)
	l.Infoln("Initial prefix")
	l.SetPrefix("customlog_upd:")
	l.SetFlags(-1)

	l.Traceln("You shouldn't see it")
	l.Infoln("Custom prefix")
	l.Warningln("Warning msg")
	l.Errorln("Error message")

	l.SetFlags(log.Ltime | log.Lshortfile)
	l.Errorln("Short time")
	l.SetLevel(LevelTrace)
	l.Traceln("Trace info")
	l.SetLevel(LevelDebug)
	l.Traceln("You shouldn't see trace")
	l.SetLevel(LevelError)
	l.Infoln("You shouldn't see info")
}
