package kata

import (
	"math/rand"
	"time"
)

type Solution struct {
	seededRand *rand.Rand
	nums       []int
}

func Constructor(nums []int) Solution {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	return Solution{
		seededRand: seededRand,
		nums:       nums,
	}
}

func (this *Solution) Pick(target int) int {

	idx := make(map[int][]int)
	for i, _ := range this.nums {
		if target == this.nums[i] {
			v, ok := idx[target]
			if !ok {
				v = make([]int, 0)
				v = append(v, i)
				idx[target] = v
				continue
			}
			v = append(v, i)
			idx[target] = v
		}
	}

	i := this.seededRand.Intn(len(idx[target]))

	return idx[target][i]
}
