// Package utils provides utility functions
package utils

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
