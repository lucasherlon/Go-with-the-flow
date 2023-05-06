// hackerranck 07
func staircase(n int32) {
    for i := 1; i <= int(n); i++ {
        for j := 1; j <= int(n); j++ {
            if i + j  < int(n) + 1 {
                fmt.Printf(" ")
            } else {
                fmt.Printf("#")
            }
        }
        fmt.Printf("\n")
    }
}