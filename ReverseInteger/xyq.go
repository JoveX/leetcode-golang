package main

import "fmt"

func main() {
    x := -123
    res := reverse(x)
    fmt.Println(res)
}

func reverse(x int) int {
    res := 0;
    for x != 0 {
        res = res * 10 + x % 10
        x = x / 10
    }

    if res < -2147483648 || res > 2147483647 {
        return 0
    }

    return res
}
