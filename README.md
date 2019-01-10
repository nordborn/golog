**Package golog is a logging package that I'd like to be initially in Go's (Golang) standard library: 
it's a simple logging tool with levels and different outputs for info and error messages.**

**TL;DR**
```Go
golog.Infoln("Started")
// Output: [INF] main: 2018/11/26 16:57:49 main.go:61: Started
golog.Errorf("An error %v occured: %v\n", mycode, mymsg)
// Output: [ERR] main: 2018/11/26 16:57:50 main.go:62: An error 01 occured: just err
```


Package "golog" provides loggerGlobal and Logger (custom logger) which support:
1. Levels:
- trace (`Trace(), Tracef(), Traceln()`);
- debug (`Debug(), Debugf(), Debugln()`);
- info (`Info(), Infof(), Infoln(), Print(), Printf(), Println()`);
- warning (`Warning(), Warningf(), Warningln()`);
- error (`Error(), Errorf(), Errorln()`);
- critical (`Critical(), Criticalf(), Criticalln()`);
- panic (`Panic(), Panicf(), Panicln()`);
- fatal (`Fatal(), Fatalf(), Fatalln()`).
2. Different outputs:
- for info-like messages (Trace, Debug, Info, Warning) - os.Stdout by default;
- for error-like messages (Error, Critical, Panic, Fatal) - os.Stderr by default.

You can set:
1. logging level `golog.SetLevel(golog.LevelInfo)` (LevelTrace, LevelDebug, LevelInfo, LevelWarning, LevelError, LevelCritical).
In this case, messages from the lower level will be omitted - LevelTrace by default;
2. custom prefix `golog.SetPrefix("myapp:")` additionally to level prefixes ("main: " by default);
3. output io.Writer interfaces `golog.SetOutput(myOutLogWriter, myErrLogWriter)`:
 - l.outWriter for Trace-Warning (os.Stdout by default);
 - l.errWriter for Error-Fatal (os.Stderr by default).
4. flags `golog.SetFlags(log.Ltime | log.Lshortfile)` similar to "log" from standard library for time and file information
("2018/11/26 16:57:49 golog.go:61" by default).

You can change level prefixes directly (defaults are TRC, DBG, INF, ERR, CRT, PNC, FTL) but don't do it
if you don't need it really.

Also, "golog" uses same position conventions as "log": all prefixes are placed before time info.

So, common message using `golog.Infoln("Started")` will be:
```
[INF] main: 2018/11/26 16:57:49 main.go:61: Started
```

Tip: in common usage if you don't know which messages you should use, use `Infoln()` and `Errorln()`.

Enjoy!
