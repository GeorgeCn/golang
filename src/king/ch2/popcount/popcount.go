package popcount

// pc[i] 是i的种群统计
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1) // i&1 与1按位与 判断奇偶性
	}
}

// PopCount 返回x的种群统计 (位置的个数)
func PopCount(x unit64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))] +
	)
}