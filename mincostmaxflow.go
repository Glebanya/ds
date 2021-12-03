package ds

import "math"

const N = 100

func min(first, second int) int {
	if first > second {
		return second
	}
	return first
}

func EdmondsKarp(c [N][N]int) int {
	f := [N][N]int{}
	for {
		from := [N]int{}
		for i := 1; i < N; i++ {
			from[i] = -1
		}
		q, h, t := [N]int{}, 0, 1
		for h < t {
			cur := q[h]
			h += 1
			for v := 0; v < N; v++ {
				if from[v] == -1 && c[cur][v]-f[cur][v] > 0 {
					q[t], from[v] = v, cur
					t += 1
				}
			}
		}
		if from[N-1] == -1 {
			break
		}

		cf := math.MaxInt64
		for cur := N - 1; cur != 0; {
			prev := from[cur]
			cf = min(cf, c[prev][cur]-f[prev][cur])
			cur = prev
		}
		for cur := N - 1; cur != 0; {
			prev := from[cur]
			f[prev][cur] += cf
			f[cur][prev] -= cf
			cur = prev
		}
	}

	flow := 0
	for i := 0; i < N; i++ {
		if c[0][i] == 0 {
			flow += f[0][i]
		}
	}

	return flow
}
