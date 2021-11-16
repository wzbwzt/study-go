package main

import "math/big"

func main() {
	//m*n的网格中，从左上角走到右下角的总方案数
	//共需要移动m+n-2次，选择n-1向右移动的总次数
}

//计算二项式
//
//m中n种组合方案
func getBio(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
