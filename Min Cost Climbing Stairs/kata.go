package kata

func minCostClimbingStairs(cost []int) int {

	dp := make([]int, len(cost))
	dp[len(cost)-1] = 0

	for i := len(cost) - 1; i >= 0; i-- {
		dp[i] = cost[i] + min(dp, i+1, i+2)
	}

	return min(dp, 0, 1)
}

func min(dp []int, i, j int) int {

	var a int
	if i < len(dp) {
		a = dp[i]
	}
	var b int
	if j < len(dp) {
		b = dp[j]
	}

	if a > b {
		return b
	}

	return a
}
