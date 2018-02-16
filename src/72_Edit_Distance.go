package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	} else if n == 0 {
		return m
	}

	sign := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		sign[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		sign[i][0] = i
	}
	for j := 0; j <= n; j++ {
		sign[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			tmp := sign[i][j-1] + 1
			if sign[i-1][j]+1 < tmp {
				tmp = sign[i-1][j] + 1
			}
			if word1[i-1] == word2[j-1] && sign[i-1][j-1] < tmp {
				tmp = sign[i-1][j-1]
			} else if sign[i-1][j-1]+1 < tmp {
				tmp = sign[i-1][j-1] + 1
			}
			sign[i][j] = tmp
		}
	}
	return sign[m][n]
}

func main() {
	fmt.Println(minDistance("ab", "a"))
	return
}
