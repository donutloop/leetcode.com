package kata

import (
	"math"
	"sort"
)

type value struct {
	v int
	d int
}

func findClosestElements(arr []int, k int, x int) []int {

	dis := make([]value, len(arr))
	for i := 0; i < len(arr); i++ {
		elem := value{
			v: arr[i],
			d: int(math.Abs(float64(arr[i]) - float64(x))),
		}
		dis[i] = elem
	}
	sort.Slice(dis, func(i, j int) bool {
		if dis[i].d == dis[j].d && dis[i].v < dis[j].v {
			return true
		} else if dis[i].d < dis[j].d {
			return true
		}
		return false
	})
	res := make([]int, k)
	for i := 0; i < len(dis[:k]); i++ {
		res[i] = dis[i].v
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}
