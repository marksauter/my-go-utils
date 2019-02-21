// Package myutil contains common utility functions I (marksauter) use across
// projects
package myutil

import (
	"fmt"
	"path/filepath"
	"reflect"
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

// Returns a string representation of value, for logging purposes, with pointers
// dereferenced if possible.
func Slog(value interface{}) string {
	if value == nil {
		return fmt.Sprintf("%v", value)
	}

	var v reflect.Value
	var ok bool
	if v, ok = value.(reflect.Value); !ok {
		v = reflect.ValueOf(value)
	}
	kind := v.Kind()
	if isArray(v) {
		ret := "["
		for i := 0; i < v.Len(); i++ {
			ret += Slog(v.Index(i))
			if i < v.Len()-1 {
				ret += " "
			}
		}
		ret += "]"
		return ret
	} else if kind == reflect.Map {
		keys := v.MapKeys()
		ret := "map["
		for i := 0; i < v.Len(); i++ {
			ret += Slog(keys[i]) + ":"
			ret += Slog(v.MapIndex(keys[i]))
			if i < v.Len()-1 {
				ret += " "
			}
		}
		ret += "]"
		return ret
	} else if kind == reflect.Struct {
		ret := "{"
		for i := 0; i < v.NumField(); i++ {
			ret += v.Type().Field(i).Name + ":"
			ret += Slog(v.Field(i))
			if i < v.NumField()-1 {
				ret += " "
			}
		}
		ret += "}"
		return ret
	} else if canBeNil(v) && v.IsNil() {
		return fmt.Sprintf("%v", v)
	} else if hasElem(v) {
		elem := v.Elem()
		return Slog(elem)
	}

	return fmt.Sprintf("%v", v)

}

func canBeNil(v reflect.Value) bool {
	k := v.Kind()
	if k == reflect.Chan ||
		k == reflect.Func ||
		k == reflect.Interface ||
		k == reflect.Map ||
		k == reflect.Ptr ||
		k == reflect.Slice {

		return true
	}
	return false
}

func isArray(v reflect.Value) bool {
	k := v.Kind()
	if k == reflect.Array || k == reflect.Slice {
		return true
	}
	return false
}

func hasElem(v reflect.Value) bool {
	k := v.Kind()
	if k == reflect.Interface || k == reflect.Ptr {
		return true
	}
	return false
}
