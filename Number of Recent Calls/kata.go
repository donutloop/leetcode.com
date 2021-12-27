package kata

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pings: make([]int, 0),
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.pings = append(rc.pings, t)
	var count int
	min := t - 3000

	i := len(rc.pings) - 1
	for ; i >= 0; i-- {
		if rc.pings[i] < min {
			break
		}
		count++
	}

	rc.pings = rc.pings[i+1:]

	return count
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
