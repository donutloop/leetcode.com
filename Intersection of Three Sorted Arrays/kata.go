package kata

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {

	res := make([]int, 0)
	var j, k int
	for i := 0; i < len(arr1); {
		if k == len(arr3) || j == len(arr2) || i == len(arr1) {
			break
		}

		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			res = append(res, arr2[j])
			j++
			i++
			k++
			continue
		}

		if arr3[k] > arr1[i] && arr3[k] > arr2[j] {
			j++
			i++
			continue
		} else if arr2[j] > arr1[i] && arr2[j] > arr3[k] {
			i++
			k++
			continue
		} else if arr1[i] > arr3[k] && arr1[i] > arr2[j] {
			j++
			k++
			continue
		} else if arr3[k] > arr1[i] {
			i++
			continue
		} else if arr3[k] > arr2[j] {
			j++
			continue
		} else if arr2[j] > arr1[i] {
			i++
			continue
		} else if arr2[j] > arr3[k] {
			k++
			continue
		} else if arr1[i] > arr3[k] {
			k++
			continue
		} else if arr1[i] > arr2[j] {
			j++
			continue
		}
	}

	return res
}
