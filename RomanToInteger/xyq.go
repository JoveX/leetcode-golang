package main

import "fmt"

func main() {
    s := "CMXC"
    res := romanToInt(s)
    fmt.Println(res)
}

func romanToInt(s string) int {
    int_map := make(map[string]int)
    int_map["I"] = 1
    int_map["V"] = 5
    int_map["X"] = 10
    int_map["C"] = 100
    int_map["M"] = 1000
    int_map["L"] = 50
    int_map["D"] = 500
    res := 0
    for i := 0; i < len(s); i++ {
        if i + 1 >= len(s) || int_map[s[i + 1:i + 2]] <= int_map[s[i:i + 1]] {
            res += int_map[s[i:i + 1]]
        } else {
            res -= int_map[s[i:i + 1]]
        }
    }
    return res
}
