package change

import (
	"errors"
)

func Change(coins []int, amount int) ([]int, error) {
	if amount < 0 || len(coins) == 0 {
		return nil, errors.New("no answer")
	} else if amount == 0 {
		return []int{}, nil
	}

	dp := make([][]int, amount+1)
	for i := range dp {
		dp[i] = make([]int, 0)
	}
	dp[0] = append(dp[0], 0)

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if len(dp[i-coin]) > 0 && (len(dp[i]) == 0 || len(dp[i-coin])+1 < len(dp[i])) {
				dp[i] = append(append([]int{}, dp[i-coin]...), coin)
			}
		}
	}

	var answer []int
	if len(dp[amount]) > 1 {
		answer = dp[amount][1:]
	} else {
		return nil, errors.New("no answer")
	}
	return answer, nil
}
