package main

import "fmt"

func main() {
    s := "aaa"
    p := "aaaa"
    res := isMatch(s, p)
    fmt.Println(res)
}

func isMatch(s string, p string) bool {
    p_len := len(p)
    s_len := len(s)

    if p_len == 0 && s_len == 0 {
        return true
    }

    if s == ".*" {
        return true
    }

    j := 0
    i := 0

    for i <= p_len {
        // 包含*
        if i < p_len - 1 && p[i+1:i+2] == "*" {
            for j < s_len - 1 && (s[j:j+1] == p[i:i+1] || p[i:i+1] == ".") {
                j++
            }
            i++
        } else {
            if j < s_len - 1 && (s[j:j+1] != p[i:i+1] || p[i:i+1] == ".") {
                return false
            }
            j++
        }
        i++
    }
    if s_len > j {
        return false;
    }
    if p_len > i {
        return false;
    }
    // pi := 0
    // si := 0
    // plen := len(p)
    // slen := len(s)


    return true
}

func maxArea1(height []int) int {
    len := len(height)
    res := 0
    for i := 0; i < len; i++ {
        for j := i + 1; j < len; j++ {
            h := 0
            if height[i] > height[j] {
                h = height[j]
            } else {
                h = height[i]
            }

            w := j - i
            if res < h * w {
                res = h * w
            }
        }
    }
    return res
}

func maxArea2(height []int) int {
    len := len(height)
    il := 0
    ir := len - 1

    res := 0
    for il < ir {
        h := 0
        if height[il] > height[ir] {
            h = height[ir]
        } else {
            h = height[il]
        }

        w := ir - il
        if res < h * w {
            res = h * w
        }

        if height[il] > height[ir] {
            ir--
        } else {
            il++
        }
    }

    return res
}
