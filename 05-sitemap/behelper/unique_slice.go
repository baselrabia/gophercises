package behelper

func SliceMoveToMap(data []string) map[string]bool {
	uniqueMap := make(map[string]bool)

	// Insert unique elements into the map
	for _, item := range data {
		uniqueMap[item] = true
	}

	return uniqueMap
}

func SliceRemoveDuplicates(data []string) []string {
	// Create a slice from the map keys
	var uniqueSlice []string
	for item := range SliceMoveToMap(data) {
		uniqueSlice = append(uniqueSlice, item)
	}

	return uniqueSlice
}
