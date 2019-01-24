// Package myutil contains common utility functions I (marksauter) use across projects
package myutil

// Returns true if the slice contains a string equal to the first argument
func ElemString(e string, ss []string) bool {
	for _, s := range ss {
		if s == e {
			return true
		}
	}
	return false
}
