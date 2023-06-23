/*
Package loggo exposes an API to log your work.

First, instantiate a logger with loggo.New, and giving it a threshold level.
Messages of lesser criticality won't be logged.

Sharing the logger is the responsibility of the caller.

The logger can be called to log messages on five levels:
  - Debug: mostly used to debug code, follow step-by-step processes
  - Info: valuable messages providing  insights to the milestones of a process
  - Warn: valuable message indicating when the application may not be having the intended outcom
  - Error: error messages to understand what went wrong
  - Fatal: error message indicating why the application crashed
*/
package loggo
