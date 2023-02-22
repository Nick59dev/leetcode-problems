package medium

func convert(s string, numRows int) string {
    if numRows <= 0{
        return ""
    }
    if numRows == 1 {
        return s
    }

    result := []rune{}
    filter := make(map[int][]rune)

    step  := -1
    point := 1

    for _, x := range s {
        if point == 1 || point == numRows {
            step *= -1
        }

        filter[point] = append(filter[point], rune(x))
        point += step
    }

    for i := 1; i <= numRows; i++ {
        for _, x := range filter[i] {
            result = append(result, rune(x))
        }
    }

    return string(result)
}
