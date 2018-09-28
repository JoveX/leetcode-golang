package main

import "fmt"
import "sort"

func main() {
    nums := []int{1, -1, -1, 0}
    res := threeSum(nums);
    fmt.Println(':', res)
}

func threeSum(nums []int) [][]int {
    sort.Ints(nums)

    len := len(nums)
    res := [][]int{}

    for i := 0; i < len - 2; i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        j := i + 1
        k := len - 1

        for j < k {
            if nums[i] + nums[j] + nums[k] < 0 {
                j++
                for j < k && nums[j - 1] == nums[j] {
                    j++
                }
            } else if nums[i] + nums[j] + nums[k] > 0 {
                k--
                for j < k && nums[k + 1] == nums[k] {
                    k--
                }
            } else {
                item := []int{nums[i], nums[j], nums[k]}
                res = append(res, item)
                j++
                k--
                for j < k && nums[j - 1] == nums[j] {
                    j++
                }
                for j < k && nums[k + 1] == nums[k] {
                    k--
                }
            }
        }
    }

    return res
}
