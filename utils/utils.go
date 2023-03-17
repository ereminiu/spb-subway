package utils

func EnToRu(s string) string {
	en := []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', 'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.'}
	ru := []rune{'й', 'ц', 'у', 'к', 'е', 'н', 'г', 'ш', 'щ', 'з', 'х', 'ф', 'ы', 'в', 'а', 'п', 'р', 'о', 'л', 'д', 'ж', 'я', 'ч', 'с', 'м', 'и', 'т', 'ь', 'б', 'ю'}

	t := ""
	for _, c := range s {
		g := c
		for i := 0; i < len(en); i++ {
			if en[i] == c {
				g = ru[i]
			} else if en[i]-32 == c {
				g = ru[i] - 32
			}
		}
		t += string(g)
	}

	return t
}

func EditDistance(s, t string) int {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}
			d := 0
			if s[i-1] != t[j-1] {
				d = 1
			}

			dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+d)
		}
	}
	return dp[n][m]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
