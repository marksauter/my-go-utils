// Package myutil contains common utility functions I (marksauter) use across
// projects
package myutil

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Returns string with function name and line number. The argument skip is the
// number of stack frames to skip, with 0 identifying the frame from which Trace
// was called, 1 identifying the frame above, etc.
func Trace(skip int) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2+skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf(
		"%s:%d:",
		filepath.Base(frame.Function),
		frame.Line,
	)
}
