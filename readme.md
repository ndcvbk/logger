Create new logger:

```
log := logger.GetInstance("LOGLEVEL")
```

Possible log levels are: trace, debug, info, warning, error, fatal

Create log entry:

```
log.Warn("Some descriptive string.")
log.Warn("Some descriptive string with parameters: %s, %d.", "param1", 2)
```
