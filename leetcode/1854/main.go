package main

/*
1854. 人口最多的年份
给你一个二维整数数组 logs ，其中每个 logs[i] = [birthi, deathi] 表示第 i 个人的出生和死亡年份。

年份 x 的 人口 定义为这一年期间活着的人的数目。第 i 个人被计入年份 x 的人口需要满足：x 在闭区间 [birthi, deathi - 1] 内。注意，人不应当计入他们死亡当年的人口中。

返回 人口最多 且 最早 的年份。
*/
func main() {
	maximumPopulation([][]int{{2008, 2026}, {2004, 2008}, {2034, 2035}, {1999, 2050}, {2049, 2050}, {2011, 2035}, {1966, 2033}, {2044, 2049}})
}

func maximumPopulation(logs [][]int) int {
	m := make(map[int]int, 0)
	for i := 0; i < len(logs); i++ {
		for j := logs[i][0]; j < logs[i][1]; j++ {
			m[j] += 1
		}
	}
	max := 0
	y := 0
	for k, v := range m {
		if v > max {
			max = v
			y = k
		} else if v == max {
			if y > k {
				y = k
			}
		}
	}
	return y
}
