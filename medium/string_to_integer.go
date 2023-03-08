package medium

import (
    "fmt"
    "math"
)

func myAtoi(s string) int {
    var sum          int = 0
    var counter      int = 0
    var factor       int = 1

    if s == "" {
        return 0
    }

    for counter != len(s) && s[counter] == ' ' {
        counter++
    }
    
    if counter < len(s) && (s[counter] == '+' || s[counter] == '-') {
        if s[counter] == '-' {
            factor = -1
        }
        counter++
    }

    for counter != len(s) && 
        (rune(s[counter]) >= 48 && rune(s[counter]) <= 57) {
            if sum > math.MaxInt32 / 10 || (sum == math.MaxInt32 / 10 && (rune(s[counter]) - 48) > 7) {
                if factor == -1 {
                    return math.MinInt32
                } else {
                    return math.MaxInt32
                }
            }

            sum = (sum * 10) + int(rune(s[counter]) - 48)
            counter++
    }

    return sum * factor
}
