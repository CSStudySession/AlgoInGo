package main

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	/*
	dp[i][j]: s的前i位和p的前j位是否match
	 */
	dp := make([][]bool, m + 1)
	for i := range dp {
		dp[i] = make([]bool, n + 1)
	}

	dp[0][0] = true // 前0位肯定match
	for i := 2; i <= n; i++ {
		if p[i - 1] == '*' {
			dp[0][i] = dp[0][i - 2] // 初始化dp矩阵的第一行
		}
	}

	// dp矩阵第一列应该都是false 而dp矩阵初始化默认值都是false 所以不需要显式初始化

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sc := s[i - 1] //第i位的下标是i-1
			pc := p[j - 1]

			if sc == pc || pc == '.' {
				dp[i][j] = dp[i - 1][j - 1]
			} else if pc == '*' {
				dp[i][j] = dp[i][j - 2] || (dp[i - 1][j] && (s[i - 1] == p[j - 2] || p[j - 2] == '.'))
			}
		}
	}
	return dp[m][n]
}