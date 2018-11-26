**Package golog is a logging package that I'd like to be initially in Go's (Golang) standard library**.

Packege golog provides globalLog and Logger which support:
1. Levels:
- trace;
- debug;
- info;
- warning;
- error;
- critical.
2. Different outputs:
- for info-like messages (Trace, Debug, Info), they use os.Stdout by default;
- for error-like messages (Warning, Error, Critical, Panic, Fatal) they use os.Stderr by default.

You can set:
1. logging level (LevelTrace, LevelDebug, LevelInfo, LevelWarning, LevelError, LevelCritical);
2. custom prefix (e.g. "[myapp]: ") additionally to level prefixes ("[main]: " by default);
3. output file-like interfaces:
 - l.outFile for Trace-Info (os.Stdout by default);
 - l.errFile for Error-Fatal (os.Stderr by default).

You can set flag similar to "log" from standard library for time and file information
(default provides "2018/11/26 16:57:49 golog.go:61").

Also, you can change level prefixes (defaults are "TRACE: ", "DEBUG: " etc.) but don't do it
if you don't need it really.

Tip: in common usage if you don't know which messages you should use, use Infoln() and Errorln().

So, common message using `golog.Infoln("Started")` will be:
```
INFO: [main]: 2018/11/26 16:57:49 main.go:61: Started
```