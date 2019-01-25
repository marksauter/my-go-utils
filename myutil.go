// Package myutil contains common utility functions I (marksauter) use across
// projects
package myutil

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Returns string with function name, line number, and message for logging
// purposes.
func Trace(message string) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf(
		"%s:%d: %s",
		filepath.Base(frame.Function),
		frame.Line,
		message,
	)
}
