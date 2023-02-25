package easy

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else if x == 0 {
        return true
    } else if x / 10 == 0 {
        return true
    }

    length := 0

    numbers := make([]byte, 0)

    for temp := x; temp != 0; {
        numbers = append(numbers, byte(temp % 10))
        length += 1
        temp /= 10
    }

    for i := 0; i < len(numbers) / 2; i++ {
        if numbers[i] != numbers[len(numbers) - 1 - i] {
            return false
        }
    }

    return true
}
