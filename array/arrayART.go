package array

// CopySlice return a copy of |arr|.
func CopySlice[T any](arr []T) []T {
	ret := make([]T, len(arr))
	copy(ret, arr)

	return ret
}

// ReverseInPlace reverses |arr| in-place.
func ReverseInPlace[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Reverse returns a new slice, with elements being those of |arr| after being reversed.
func Reverse[T any](arr []T) []T {
	reversed := CopySlice(arr)
	ReverseInPlace(reversed)

	return reversed
}

//============================================================================================
