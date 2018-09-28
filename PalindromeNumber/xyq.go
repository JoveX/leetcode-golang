package main

import "fmt"

func main() {
    x := 121
    res := isPalindrome(x)
    fmt.Println(res)
}
func isPalindrome(x int) bool {
    if x <= 9 && x >= 0 {
        return true
    }

    if x < 0 || x % 10 == 0 {
        return false
    }

    reverse := 0
    num := x
    for x != 0 {
        reverse = reverse * 10 + x % 10
        x /= 10
    }

    return reverse == num
}
