package main

import (
	"fmt"
	"strings"
	"strconv"
)

func Is_valid_ip(ip string) bool {
	for _, char := range ip {
		if (int(char) > 57 || int(char) < 48) && int(char) != 46 {
			return false
		}
	}
	list := strings.Split(ip, ".")
	if len(list) != 4 {
		return false
	}

	for _, val := range list {
		if len(val) > 3 {
			return false
		}

		if len(val) > 1 && val[0] == '0' {
			return false
		}

		num, _ := strconv.Atoi(val)
		if num > 255 || num < 0 {
			return false
		}

	}

	return true
}

func main() {
	example := "0.34.82.5"
	res := Is_valid_ip(example)
	fmt.Println(res)
}