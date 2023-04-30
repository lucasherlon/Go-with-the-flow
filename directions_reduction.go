package main

import "fmt"

func count(dirs []string, dir string) int {
	counter := 0
	for _, aim := range dirs {
		if aim == dir {
			counter++
		}
	}
	return counter
}

func nulify(dirs []string, dir string, quant int) []string {
	counter := 0
	for i, aim := range dirs {
		if aim == dir {
			dirs[i] = "0"
			counter++
		}
		if counter == quant {
			break
		}
	}
	return dirs
}
 
func DirReduc(arr []string) []string {
    var newArr []string

   	north := count(arr, "NORTH")
   	south := count(arr, "SOUTH")
   	east := count(arr, "EAST")
   	west := count(arr, "WEST")

   	if north >= south {
   		nulify(arr, "SOUTH", south)
   		nulify(arr, "NORTH", south)
   	} else {
   		nulify(arr, "SOUTH", north)
   		nulify(arr, "NORTH", north)
   	}

   	if east >= west {
   		nulify(arr, "EAST", west)
   		nulify(arr, "WEST", west)
   	} else {
   		nulify(arr, "EAST", east)
   		nulify(arr, "WEST", east)
   	}

	for _, dir := range arr {
		if dir != "0" {
			newArr = append(newArr, dir)
		}
   	}
   	
  	return newArr

}

func main() {
	dirs := []string{"NORTH", "SOUTH", "EAST", "WEST"}
	newDirs := DirReduc(dirs)
	fmt.Println(newDirs)
}