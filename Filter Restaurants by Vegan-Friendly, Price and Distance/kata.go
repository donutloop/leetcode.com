package kata

import "sort"

const (
	idIndex            int = 0
	ratingIndex        int = 1
	veganFriendlyIndex int = 2
	priceIndex         int = 3
	distanceIndex      int = 4
)

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {

	matchedRestaurant := make([][]int, 0)
	for i := range restaurants {
		if veganFriendly != 0 {
			if restaurants[i][veganFriendlyIndex] != veganFriendly {
				continue
			}
		}
		if maxPrice < restaurants[i][priceIndex] {
			continue
		}
		if maxDistance < restaurants[i][distanceIndex] {
			continue
		}
		matchedRestaurant = append(matchedRestaurant, restaurants[i])
	}

	sort.Slice(matchedRestaurant, func(i, j int) bool {
		if matchedRestaurant[i][ratingIndex] == matchedRestaurant[j][ratingIndex] {
			return matchedRestaurant[i][idIndex] > matchedRestaurant[j][idIndex]
		}
		return matchedRestaurant[i][ratingIndex] > matchedRestaurant[j][ratingIndex]
	})

	ids := make([]int, len(matchedRestaurant))
	for i, restaurant := range matchedRestaurant {
		ids[i] = restaurant[idIndex]
	}

	return ids
}
