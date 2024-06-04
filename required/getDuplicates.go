package required

func GetDuplicates(x []int) []int {
	var duplicate []int

	for i := 0; i < len(x); i++ {
		isDuplicate := false
		for j := 0; j < len(duplicate); j++ {
			if x[i] == duplicate[j] {
				isDuplicate = true
				break
			}
		}

		if isDuplicate {
			continue
		}

		for j := i + 1; j < len(x); j++ {
			if x[i] == x[j] {
				duplicate = append(duplicate, x[i])
				break
			}
		}
	}

	if len(duplicate) > 0 {
		return duplicate
	} else {
		return nil
	}
}
