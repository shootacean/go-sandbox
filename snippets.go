// go-snippets is snippets with golang
package go_snippets

import (
	"math"
)

// ある要素を削除したスライスを返す
func remove(nums []int, i int) []int {
	return append(nums[:i], nums[i+1:]...)
}

// スライス内の最大値を取得する
func max(nums []int) (index, max int) {
	max = math.MinInt64
	for i, n := range nums {
		if n > max {
			index = i
			max = n
		}
	}
	return
}
