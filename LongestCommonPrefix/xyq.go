package main

import "fmt"

func main() {
    strs := []string{"a", "b"}
    res := longestCommonPrefix(strs)
    fmt.Println(res)
}

func longestCommonPrefix(strs []string) string {
    if len(strs) < 1 {
        return ""
    }
    max_prefix, strs := strs[0], strs[1:]
    max_len := len(max_prefix)

    for _, str := range strs {
        if max_len < 1 {
            return ""
        }
        if max_len > len(str) {
            max_len = len(str)
        }
        for max_len >= 1 {
            if max_prefix[0:max_len] == str[0:max_len] {
                max_prefix = max_prefix[0:max_len]
                break
            }
            max_len--
        }
    }
    if max_len == 0 {
        return ""
    }

    return max_prefix
}
