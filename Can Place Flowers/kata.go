package kata

func canPlaceFlowers(flowerbed []int, n int) bool {

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if i-1 >= 0 && flowerbed[i-1] != 0 {
				continue
			}
			if i+1 < len(flowerbed) && flowerbed[i+1] != 0 {
				continue
			}
			flowerbed[i] = 1
			n = n - 1
		}
	}
	return n <= 0
}
