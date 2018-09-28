package main

import "fmt"

func main() {
    nums := []int{1, 2, 4, 3}
    res := maxArea2(nums)
    fmt.Println(res)
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
