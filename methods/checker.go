package methods

func checkIfExists(arr []string, search string) bool {
	for _, item := range arr {
		if item == search {
			return true
		}
	}
	return false
}
