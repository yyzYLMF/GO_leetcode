package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	if matrix[0][0] > target {
		return false
	}

	low, high := 0, m-1
	for {
		if low >= high {
			break
		}
		mid := (low + high) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	pos := low
	if matrix[pos][0] > target {
		pos = pos - 1
	}

	if matrix[pos][n-1] < target {
		return false
	}
	low, high = 0, n-1
	for {
		if low > high {
			break
		}
		mid := (low + high) / 2
		if matrix[pos][mid] == target {
			return true
		} else if matrix[pos][mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	return
}
