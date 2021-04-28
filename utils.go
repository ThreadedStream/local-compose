package main

func find(arr []vertex, name string) bool {
	//TODO: Improve complexity. Runs in linear time now
	//Analysis: O(n) - Worst case, O(1) best case,
	for _, value := range arr {
		if value.name == name {
			return true
		}
	}
	return false
}

func findService(name string) bool {
	for _, value := range services {
		if value == name {
			return true
		}
	}
	return false
}
