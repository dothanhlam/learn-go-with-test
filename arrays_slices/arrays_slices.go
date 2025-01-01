package arrays_slices

import (
    "unicode"
)

// Sum calculates the sum of a slice of integers.
func Sum(numbers []int) int {
    return sum(numbers)
}

// SumWithRange calculates the sum of a slice of integers using a range loop.
func SumWithRange(numbers []int) int {
    return sum(numbers)
}

// SumDigits calculates the sum of all digit characters in a string.
func SumDigits(input string) int {
    sum := 0
    for _, char := range input {
        if unicode.IsDigit(char) {
            sum += int(char - '0')
        }
    }
    return sum
}



// sum is a helper function to calculate the sum of a slice of integers.
func sum(numbers []int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}