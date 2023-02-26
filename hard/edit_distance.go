package hard


func min(numbers ...int) (minimum int) {
    if len(numbers) == 1 {
        return numbers[0]
    }
    minimum = numbers[0]

    for _, v := range numbers {
        if v < minimum {
            minimum = v
        }
    }

    return
}

func minDistance(word1 string, word2 string) int {
    return minDistanceMain(word1, word2, len(word1), len(word2))
}

func minDistanceMain(word1, word2 string, word1Index, word2Index int) int {
    if word1Index == 0 {
        return word2Index
    }
    if word2Index == 0 {
        return word1Index
    }
    if word1[word1Index - 1] == word2[word2Index - 1] {
        return minDistanceMain(word1, word2, word1Index - 1, word2Index - 1)
    } else {
        insertOp := minDistanceMain(word1, word2, word1Index, word2Index - 1)
        deleteOp := minDistanceMain(word1, word2, word1Index - 1, word2Index)
        replaceOp := minDistanceMain(word1, word2, word1Index - 1, word2Index - 1)

        return min(insertOp, min(deleteOp, replaceOp)) + 1
    }
}
