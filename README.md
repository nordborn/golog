**Package golog is a logging package that I'd like to be initially in Go's (Golang) standard library: 
it's a simple logging tool with levels and different outputs for info and error messages.**

So, you can use "golog" for all  

Package golog provides globalLog and Logger which support:
1. Levels:
- trace (Trace(), Tracef(), Traceln());
- debug (Debug(), Debugf(), Debugln());
- info (Info(), Infof(), Infoln(), Print(), Printf(), Println();
- warning (Warning(), Warningf(), Warningln());
- error (Error(), Errorf(), Errorln());
- critical (Critical(), Criticalf(), Criticalln());
- panic (Panic(), Panicf(), Panicln());
- fatal (Fatal(), Fatalf(), Fatalln()).
2. Different outputs:
- for info-like messages (Trace, Debug, Info) - os.Stdout by default;
- for error-like messages (Warning, Error, Critical, Panic, Fatal) - os.Stderr by default.

You can set:
1. logging level (LevelTrace, LevelDebug, LevelInfo, LevelWarning, LevelError, LevelCritical).
In this case, messages from the lower level will be omitted - LevelTrace by default;
2. custom prefix (e.g. "myapp: ") additionally to level prefixes ("main: " by default);
3. output io.Writer interfaces:
 - l.outFile for Trace-Info (os.Stdout by default);
 - l.errFile for Warning-Fatal (os.Stderr by default).

You can use `SetFlags()` similar to "log" from standard library for time and file information
("2018/11/26 16:57:49 golog.go:61" by default).

You can change level prefixes (defaults are TRC, DBG, INF, ERR, CRT, PNC, FTL) but don't do it
if you don't need it really.

Also, "golog" uses same position conventions as "log": all prefixes are placed before time info.

So, common message using `golog.Infoln("Started")` will be:
```
[INF] main: 2018/11/26 16:57:49 main.go:61: Started
```

Tip: in common usage if you don't know which messages you should use, use Infoln() and Errorln().

Enjoy!
