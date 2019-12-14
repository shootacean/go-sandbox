package compute

// 階乗を計算する
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 組み合わせを計算する
func combination(n, r int) int {
	return (factorial(n) / factorial(n - r)) / factorial(r)
}
