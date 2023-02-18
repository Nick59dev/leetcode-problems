package hard


func firstMissingPositive(nums []int) int {
    filter := make(map[int]bool)

    var maximum int = 0

    for _, x := range nums {
        if x > maximum {
            maximum = x
        }
        filter[x] = true
    }
    for i := 1; i <= maximum; i++ {
        if filter[i] == false {
            return i
        }
    }
    return maximum + 1
}
