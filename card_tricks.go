package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index > len(slice) - 1 {
        return -1
    }
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index > len(slice) - 1 {
        slice = append(slice, value)
        return slice
    }
    slice[index] = value
    return slice
}


// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if len(values) == 0 {
        return slice
    } else if len(values) == 1 {
    	value := values[0]
        slice = append([]int{value}, slice...)
        return slice
    } else {
    	slice = append(values, slice...)
    	return slice
    }

}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice) - 1 {
        return slice
    } else if index == 0 {
    	return slice[1:]
    } else if index == len(slice) - 1 {
    	ind := len(slice) - 1
        return slice[:ind]
    } else {
    	slice = append(slice[:index], slice[index + 1:]...)
        return slice
    }

}
