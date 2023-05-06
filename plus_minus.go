// hackerranck 06
func plusMinus(arr []int32) {
    var pos int
    var neg int
    var zero int
    for _, val := range arr {
        if val > 0 {
            pos++
        } else if val < 0 {
            neg++
        } else {
            zero++
        }
    }
    posRatio := float32(pos) / float32(len(arr))
    negRatio := float32(neg) / float32(len(arr))
    zeroRatio := float32(zero) / float32(len(arr))
    fmt.Printf("%.6f\n%.6f\n%.6f\n", posRatio, negRatio, zeroRatio)
}