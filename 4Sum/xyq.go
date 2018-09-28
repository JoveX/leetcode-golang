package main

import "fmt"
import "sort"

func main() {
    nums := []int{-2,0,0,3,3,-1}
    target := 5
    res := fourSum(nums, target);
    fmt.Println(':', res)
}

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)

    len := len(nums)
    res := [][]int{}

    for m := 0; m < len - 3; m++ {
        if m > 0 && nums[m] == nums[m - 1] {
            continue
        }

        for i := m + 1; i < len - 2; i++ {
            if i > m + 1 && nums[i] == nums[i - 1] {
                continue
            }
            j := i + 1
            k := len - 1

            for j < k {
                if nums[m] + nums[i] + nums[j] + nums[k] < target {
                    j++
                    for j < k && nums[j - 1] == nums[j] {
                        j++
                    }
                } else if nums[m] + nums[i] + nums[j] + nums[k] > target {
                    k--
                    for j < k && nums[k + 1] == nums[k] {
                        k--
                    }
                } else {
                    item := []int{nums[m], nums[i], nums[j], nums[k]}
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
    }

    return res
}
