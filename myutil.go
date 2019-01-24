// Package myutil contains common utility functions I (marksauter) use across projects
package myutil

// Returns a pointer to a string
func NewString(s string) *string {
	return &s
}

// Returns true if the slice contains a string equal to the first argument
func ElemString(e string, ss []string) bool {
	for _, s := range ss {
		if s == e {
			return true
		}
	}
	return false
}

// Returns optional head string of a []string
func HeadStringMay(ss []string) *string {
	if len(ss) > 0 {
		return &ss[0]
	}
	return nil
}
