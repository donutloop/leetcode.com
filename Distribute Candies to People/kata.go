package kata

func distributeCandies(candies int, num_people int) []int {
	people := make([]int, num_people)
	cd := 1
	var i int
	for candies > 0 {
		if i == len(people) {
			i = 0
		}

		if candies <= cd {
			cd = candies
		}

		people[i] += cd
		candies = candies - cd
		cd++
		i++
	}

	return people
}
