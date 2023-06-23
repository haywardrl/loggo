package loggo

// Level represents the logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information that is deemed valuable.
	LevelInfo
	// LevelWarn represents a logging level that contains useful information that highlights where the application has perhaps strayed from its desired functionality.
	LevelWarn
	// LevelError represents a logging level used to locate errors that have occurred within code.
	LevelError
	// LevelFatal represents the highest level of logging and is used to diagnose why an application has crashed.
	LevelFatal
)
