package main

import "fmt"

func main() {
    n := 1
    res := generateParenthesis(n)
    fmt.Println(res)
}

var res []string
func generateParenthesis(n int) []string {
    res = []string{}
    iteration("", n, n)

    return res
}

func iteration(ans string, left int, right int) {
    if left == 0 && right == 0 {
        res = append(res, ans)
    }
    if left > 0 {
        iteration(ans + "(", left - 1, right)
    }
    if right > left {
        iteration(ans + ")", left, right - 1)
    }
}
