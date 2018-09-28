package main

import "fmt"

func main() {
    s := "AB"
    numRows := 1
    res := convert(s, numRows)
    fmt.Println(res)
}

func convert(s string, numRows int) string {
    s_len := len(s)
    if s_len <= numRows || numRows == 1 {
        return s
    }
    arr := make(map[int][]string)
    m := 0
    flag := 1
    for i := 0; i < s_len; i++ {

        if _, ok := arr[m]; !ok {
            arr[m] = []string{}
        }

        arr[m] = append(arr[m], s[i:i+1])

        m += flag

        if m >= numRows - 1 {
            flag = -1
        } else if m < 1 {
            flag = 1
        }
    }

    res := ""
    for j := 0; j < numRows; j++ {
        if _, ok := arr[j]; !ok {
            return res
        }
        for k := 0; k < len(arr[j]); k++ {
            res += arr[j][k]
        }
    }

    return res
}
