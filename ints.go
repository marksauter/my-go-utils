// ints.go includes utility functions for use with integers
// the functions include variants for int32 and int64 types
package myutil

import "fmt"

// Returns a pointer to an int
func JustI(v int) *int {
	return &v
}

// Returns the value of an int pointer, or panics if the pointer is nil.
func FromJustI(v *int) int {
	if v == nil {
		panic("myutil.FromJustI: nil pointer")
	}
	return *v
}

// Returns the value of an int pointer, or 'defaultValue' if the pointer is nil.
func FromMaybeI(defaultValue int, v *int) int {
	if v == nil {
		return defaultValue
	}
	return *v
}

// Returns true if both int pointers are the same or have the same int
// value, false otherwise.
func EqI(a, b *int) bool {
	if a == b {
		return true
	} else if a != nil && b != nil {
		return *a == *b
	} else {
		return false
	}
}

// Returns a copy of a sring pointer value, or nil if the pointer is nil.
func CopyI(v *int) *int {
	if v == nil {
		return nil
	}
	cp := *v
	return &cp
}

// Returns a string representation of the int pointer: if nil then "<nil>", else
// the value of the int.
func SprintI(v *int) string {
	if v == nil {
		return fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("%d", *v)
}

// Returns a pointer to the first int in the slice, or nil if the slice is
// empty
func HeadMayI(is []int) *int {
	if len(is) > 0 {
		return &is[0]
	}
	return nil
}

// Returns a pointer to the last int in the slice, or nil if the slice is
// empty
func TailMayI(is []int) *int {
	if len(is) > 0 {
		return &is[len(is)-1]
	}
	return nil
}

// Returns the first index of the target int `t`, or
// -1 if no match is found.
func IndexI(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Returns `true` if the target int t is in the slice.
func HasI(vs []int, t int) bool {
	return IndexI(vs, t) >= 0
}

// Returns `true` if one of the ints in the slice
// satisfies the predicate `f`.
func AnyI(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// Returns `true` if all of the ints in the slice
// satisfy the predicate `f`.
func AllI(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Returns a new slice containing all ints in the
// slice that satisfy the predicate `f`.
func FilterI(vs []int, f func(int) bool) []int {
	vsf := make([]int, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Returns a new slice containing the results of applying
// the function `f` to each int in the original slice.
func MapI(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Returns a new int pointer with `f` applied to the value at `v` if not nil.
func ApplyI(v *int, f func(int) int) *int {
	if v != nil {
		va := f(*v)
		return &va
	}
	return v
}

// Returns a pointer to an int32
func JustI32(v int32) *int32 {
	return &v
}

// Returns the value of an int32 pointer, or panics if the pointer is nil.
func FromJustI32(v *int32) int32 {
	if v == nil {
		panic("myutil.FromJustI32: nil pointer")
	}
	return *v
}

// Returns the value of an int32 pointer, or 'defaultValue' if the pointer is nil.
func FromMaybeI32(defaultValue int32, v *int32) int32 {
	if v == nil {
		return defaultValue
	}
	return *v
}

// Returns true if both int32 pointers are the same or have the same int32
// value, false otherwise.
func EqI32(a, b *int32) bool {
	if a == b {
		return true
	} else if a != nil && b != nil {
		return *a == *b
	} else {
		return false
	}
}

// Returns a copy of a sring pointer value, or nil if the pointer is nil.
func CopyI32(v *int32) *int32 {
	if v == nil {
		return nil
	}
	cp := *v
	return &cp
}

// Returns a string representation of the int32 pointer: if nil then "<nil>", else
// the value of the int32.
func SprintI32(v *int32) string {
	if v == nil {
		return fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("%d", *v)
}

// Returns a pointer to the first int32 in the slice, or nil if the slice is
// empty
func HeadMayI32(is []int32) *int32 {
	if len(is) > 0 {
		return &is[0]
	}
	return nil
}

// Returns a pointer to the last int32 in the slice, or nil if the slice is
// empty
func TailMayI32(is []int32) *int32 {
	if len(is) > 0 {
		return &is[len(is)-1]
	}
	return nil
}

// Returns the first index of the target int32 `t`, or
// -1 if no match is found.
func I32ndexI32(vs []int32, t int32) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Returns `true` if the target int32 t is in the slice.
func HasI32(vs []int32, t int32) bool {
	return I32ndexI32(vs, t) >= 0
}

// Returns `true` if one of the int32s in the slice
// satisfies the predicate `f`.
func AnyI32(vs []int32, f func(int32) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// Returns `true` if all of the int32s in the slice
// satisfy the predicate `f`.
func AllI32(vs []int32, f func(int32) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Returns a new slice containing all int32s in the
// slice that satisfy the predicate `f`.
func FilterI32(vs []int32, f func(int32) bool) []int32 {
	vsf := make([]int32, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Returns a new slice containing the results of applying
// the function `f` to each int32 in the original slice.
func MapI32(vs []int32, f func(int32) int32) []int32 {
	vsm := make([]int32, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Returns a new int32 pointer with `f` applied to the value at `v` if not nil.
func ApplyI32(v *int32, f func(int32) int32) *int32 {
	if v != nil {
		va := f(*v)
		return &va
	}
	return v
}

// Returns a pointer to an int64
func JustI64(v int64) *int64 {
	return &v
}

// Returns the value of an int64 pointer, or panics if the pointer is nil.
func FromJustI64(v *int64) int64 {
	if v == nil {
		panic("myutil.FromJustI64: nil pointer")
	}
	return *v
}

// Returns the value of an int64 pointer, or 'defaultValue' if the pointer is nil.
func FromMaybeI64(defaultValue int64, v *int64) int64 {
	if v == nil {
		return defaultValue
	}
	return *v
}

// Returns true if both int64 pointers are the same or have the same int64
// value, false otherwise.
func EqI64(a, b *int64) bool {
	if a == b {
		return true
	} else if a != nil && b != nil {
		return *a == *b
	} else {
		return false
	}
}

// Returns a copy of a sring pointer value, or nil if the pointer is nil.
func CopyI64(v *int64) *int64 {
	if v == nil {
		return nil
	}
	cp := *v
	return &cp
}

// Returns a string representation of the int64 pointer: if nil then "<nil>", else
// the value of the int64.
func SprintI64(v *int64) string {
	if v == nil {
		return fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("%d", *v)
}

// Returns a pointer to the first int64 in the slice, or nil if the slice is
// empty
func HeadMayI64(is []int64) *int64 {
	if len(is) > 0 {
		return &is[0]
	}
	return nil
}

// Returns a pointer to the last int64 in the slice, or nil if the slice is
// empty
func TailMayI64(is []int64) *int64 {
	if len(is) > 0 {
		return &is[len(is)-1]
	}
	return nil
}

// Returns the first index of the target int64 `t`, or
// -1 if no match is found.
func I64ndexI64(vs []int64, t int64) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Returns `true` if the target int64 t is in the slice.
func HasI64(vs []int64, t int64) bool {
	return I64ndexI64(vs, t) >= 0
}

// Returns `true` if one of the int64s in the slice
// satisfies the predicate `f`.
func AnyI64(vs []int64, f func(int64) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// Returns `true` if all of the int64s in the slice
// satisfy the predicate `f`.
func AllI64(vs []int64, f func(int64) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Returns a new slice containing all int64s in the
// slice that satisfy the predicate `f`.
func FilterI64(vs []int64, f func(int64) bool) []int64 {
	vsf := make([]int64, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Returns a new slice containing the results of applying
// the function `f` to each int64 in the original slice.
func MapI64(vs []int64, f func(int64) int64) []int64 {
	vsm := make([]int64, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Returns a new int64 pointer with `f` applied to the value at `v` if not nil.
func ApplyI64(v *int64, f func(int64) int64) *int64 {
	if v != nil {
		va := f(*v)
		return &va
	}
	return v
}
