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

// String implements the fmt.Stringer interface
func (level Level) String() string {
	switch level {
	case LevelDebug:
		return "[DEBUG]"
	case LevelInfo:
		return "[INFO]"
	case LevelWarn:
		return "[WARN]"
	case LevelError:
		return "[ERROR]"
	case LevelFatal:
		return "[FATAL]"
	default:
		// Should not happen.
		return ""
	}
}
