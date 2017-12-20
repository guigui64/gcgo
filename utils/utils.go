// Package utils provides utility functions
package utils

import (
	"fmt"
	"log"
)

// Check checks if the error is nil and panic otherwise
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// CopyIntSlice returns a new slice as a copy of the given origin
func CopyIntSlice(origin []int) []int {
	cop := make([]int, len(origin))
	copy(cop, origin)
	return cop
}

// CopyByteSlice returns a new slice as a copy of the given origin
func CopyByteSlice(origin []byte) []byte {
	cop := make([]byte, len(origin))
	copy(cop, origin)
	return cop
}

// MinMaxIntSlice returns the min and max values and their indexes
func MinMaxIntSlice(s []int) (imin, vmin, imax, vmax int) {
	vmax, vmin = s[0], s[0]
	for i, v := range s {
		if v > vmax {
			imax, vmax = i, v
		}
		if v < vmin {
			imin, vmin = i, v
		}
	}
	return
}

// MinMaxByteSlice returns the min and max values and their indexes
func MinMaxByteSlice(s []byte) (imin int, vmin byte, imax int, vmax byte) {
	vmax, vmin = s[0], s[0]
	for i, v := range s {
		if v > vmax {
			imax, vmax = i, v
		}
		if v < vmin {
			imin, vmin = i, v
		}
	}
	return
}

func MinMax(s interface{}) (imin, imax int) {
	switch v := s.(type) {
	case []int:
		imin, _, imax, _ = MinMaxIntSlice(v)
	case []byte:
		imin, _, imax, _ = MinMaxByteSlice(v)
	default:
		log.Fatal(fmt.Sprint("%T not hanlded by utils.MinMax", v))
	}
	return
}

// MaxIntSlice looks for the max value of the given slice and returns the index
// of the first element in the slice with the max value and the value itself
func MaxIntSlice(s []int) (imax, vmax int) {
	vmax = s[0]
	for i, v := range s {
		if v > vmax {
			imax = i
			vmax = v
		}
	}
	return
}

// MaxByteSlice looks for the max value of the given slice and returns the index
// of the first element in the slice with the max value and the value itself
func MaxByteSlice(s []byte) (imax int, vmax byte) {
	vmax = s[0]
	for i, v := range s {
		if v > vmax {
			imax = i
			vmax = v
		}
	}
	return
}

// Max returns the index of the element with the higher value
func Max(s interface{}) (imax int) {
	switch v := s.(type) {
	case []int:
		imax, _ = MaxIntSlice(v)
	case []byte:
		imax, _ = MaxByteSlice(v)
	default:
		log.Fatal(fmt.Sprint("%T not hanlded by utils.Max", v))
	}
	return
}

// MinIntSlice looks for the min value of the given slice and returns the index
// of the first element in the slice with the min value and the value itself
func MinIntSlice(s []int) (imin, vmin int) {
	vmin = s[0]
	for i, v := range s {
		if v < vmin {
			imin = i
			vmin = v
		}
	}
	return
}

// MinByteSlice looks for the min value of the given slice and returns the index
// of the first element in the slice with the min value and the value itself
func MinByteSlice(s []byte) (imin int, vmin byte) {
	vmin = s[0]
	for i, v := range s {
		if v < vmin {
			imin = i
			vmin = v
		}
	}
	return
}

// Min returns the index of the element with the higher value
func Min(s interface{}) (imin int) {
	switch v := s.(type) {
	case []int:
		imin, _ = MinIntSlice(v)
	case []byte:
		imin, _ = MinByteSlice(v)
	default:
		log.Fatal(fmt.Sprint("%T not hanlded by utils.Min", v))
	}
	return
}

// InsertInt inserts an int in a slice at a specified position
func InsertInt(s []int, pos, value int) []int {
	s = append(s, 0)
	copy(s[pos+1:], s[pos:])
	s[pos] = value
	return s
}

// DeleteInt deletes element at specified position WITHOUT preserving order
func DeleteInt(s []int, pos int) []int {
	s[pos] = s[len(s)-1]
	s = s[:len(s)-1]
	return s
}
