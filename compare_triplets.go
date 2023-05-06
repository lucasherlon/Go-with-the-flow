
// hackerranck 03
func compareTriplets(a []int32, b []int32) []int32 {
    points := []int32{0, 0}
    for i := 0; i < 3; i++ {
        if a[i] > b[i] {
            points[0]++
        } else if a[i] < b[i] {
            points[1]++
        }
    }
    return points
}
