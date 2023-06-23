package loggo_test

import "github.com/haywardrl/loggo"

func ExampleLogger_Debugf() {
	debugLogger := loggo.New(loggo.LevelDebug)
	debugLogger.Errorf("This is a debugging error for %s", "logger_test")
	// Output: [ERROR] This is a debugging error for logger_test
}

func ExampleLogger_Fatalf() {
	debugLogger := loggo.New(loggo.LevelFatal)
	debugLogger.Errorf("This is a debugging error for %s", "logger_test")
	// Output:
}
