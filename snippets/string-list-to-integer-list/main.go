package main

import "strconv"


// The StringListToIntegerList function simply gets a list
// of Strings that contain integer values and  returns a list 
// of integers based on the given list values. In case a value
// is not convertable for any reasons, it will return an error.
func StringListToIntegerList(l []string) ([]int, error) {
	
	var lInt = make([]int, len(l))

	for idx, i := range l {
		j, err := strconv.Atoi(i)
		if err != nil {
			return lInt, err
		}
		lInt[idx] = j
	}

	return lInt, nil
}
