package main

import "fmt"

func main() {
    s := "abcabcbb"
    res := lengthOfLongestSubstring(s);
    fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
    s_len := len(s)
    s_map := make(map[string]int)
    ans := 0
    i := 0
    for j := 0; j < s_len; j++ {
        if _, ok := s_map[s[j:j+1]]; ok {
            if s_map[s[j:j+1]] > i {
                i = s_map[s[j:j+1]]
            }
        }
        s_map[s[j:j+1]] = j + 1

        if j - i + 1 > ans {
            ans = j - i + 1
        }
    }
    return ans
}
