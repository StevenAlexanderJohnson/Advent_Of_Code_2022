package Search

type NotFoundError struct{}

func (m *NotFoundError) Error() string {
	return "Not Found"
}

func Binary_Search(input_string []byte, target byte) (int, error) {
	var min int = 0
	var max int = len(input_string)
	var mid int = (min + max) / 2
	for mid < max {
		if input_string[mid] == target {
			return mid, nil
		} else if input_string[mid] > target {
			min = mid
			// dividing rounds down, so if the mid is the same as the previous mid, we need to increment mid
			temp_mid := (min + max) / 2
			if temp_mid == mid {
				mid++
				continue
			}
			// otherwise, we can just set mid to the new value
			mid = temp_mid
		} else if input_string[mid] < target {
			max = mid
			mid = (min + max) / 2
		}
	}
	return 0, &NotFoundError{}
}

func Count_Instances(input_string []byte, target byte) int {
	var count int = 0
	for _, char := range input_string {
		if char == target {
			count++
		}
	}
	return count
}
