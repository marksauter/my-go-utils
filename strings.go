package myutil

// Returns a pointer to a string
func NewS(s string) *string {
	return &s
}

// Returns a pointer to the first string in the slice, or nil if the slice is
// empty
func HeadMayS(ss []string) *string {
	if len(ss) > 0 {
		return &ss[0]
	}
	return nil
}

// Returns a pointer to the last string in the slice, or nil if the slice is
// empty
func TailMayS(ss []string) *string {
	if len(ss) > 0 {
		return &ss[len(ss)-1]
	}
	return nil
}

// Returns the value of a string pointer, or empty string if the pointer is nil
func JustS(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Returns the first index of the target string `t`, or
// -1 if no match is found.
func IndexS(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Returns `true` if the target string t is in the slice.
func HasS(vs []string, t string) bool {
	return IndexS(vs, t) >= 0
}

// Returns `true` if one of the strings in the slice
// satisfies the predicate `f`.
func AnyS(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// Returns `true` if all of the strings in the slice
// satisfy the predicate `f`.
func AllS(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Returns a new slice containing all strings in the
// slice that satisfy the predicate `f`.
func FilterS(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Returns a new slice containing the results of applying
// the function `f` to each string in the original slice.
func MapS(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Returns a new string pointer with `f` applied to the value at `v` if not nil.
func ApplyS(v *string, f func(string) string) *string {
	if v != nil {
		va := f(*v)
		return &va
	}
	return v
}
