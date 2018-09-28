package main

import "fmt"

func main() {
    num := 990
    res := intToRoman(num)
    fmt.Println(res)
}


func intToRoman(num int) string {
    romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

    if num <= 0 {
        return ""
    }
    res := ""
    for i := 0; i < 13 && num > 0; i++ {
        if num < nums[i] {
            continue
        }
        for num >= nums[i] {
            num -= nums[i]
            res += romans[i]
        }

    }
    return res
}
