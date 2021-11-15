package ds

import (
	"errors"
	"math"
)

func FibonacciRecursive(n int) (float64, error) {

	if n < 0 {
		return 0.0, errors.New("wrong input number")
	} else if n == 1 || n == 2 {
		return 1, nil
	} else {

		first, err := FibonacciRecursive(n - 1)
		if err != nil {
			return 0., err
		}

		second, err := FibonacciRecursive(n - 1)

		if err != nil {
			return 0., err
		}

		return first + second, nil
	}
}

func FibonacciDp(n int) (float64, error) {

	if n < 0 {
		return 0.0, errors.New("wrong input number")
	} else if n == 1 || n == 2 {
		return 1, nil
	} else {
		prev, current := 1., 1.
		for i := 0; i < n; i++ {
			current, prev = current+prev, current
		}
		return current, nil
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//link https://ru.wikipedia.org/wiki/%D0%97%D0%B0%D0%B4%D0%B0%D1%87%D0%B0_%D0%BE_%D0%BF%D0%BE%D1%80%D1%8F%D0%B4%D0%BA%D0%B5_%D0%BF%D0%B5%D1%80%D0%B5%D0%BC%D0%BD%D0%BE%D0%B6%D0%B5%D0%BD%D0%B8%D1%8F_%D0%BC%D0%B0%D1%82%D1%80%D0%B8%D1%86
func MatrixChain(matrixChain []int) int {
	n := len(matrixChain) - 1
	dp := make([]int, n*n)

	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l
			dp[i*n+j] = math.MaxInt32
			for k := i; k < j; k++ {
				dp[i*n+j] = Min(dp[i*n+j], dp[i*n+k]+dp[(k+1)*n+j]+matrixChain[i]*matrixChain[k+1]*matrixChain[j+1])
			}
		}
	}

	return dp[n-1]
}
