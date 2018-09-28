package main

import "fmt"
import "sort"

func main() {
    nums := []int{1,1,1,0}
    target := 100
    res := threeSumClosest(nums, target)
    fmt.Println(':', res)
}

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)

    len := len(nums)
    res := -1
    closest_num := -1

    for i := 0; i < len - 2; i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        j := i + 1
        k := len - 1

        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum == target {
                return target
            }

            num := target - sum
            if num < 0 {
                num = -num
            }

            if res == -1 || closest_num == -1 || num < closest_num {
                res = sum
                closest_num = num
            }

            if sum < target {
                j++
                for j < k && nums[j - 1] == nums[j] {
                    j++
                }
            } else if sum > target {
                k--
                for j < k && nums[k + 1] == nums[k] {
                    k--
                }
            }
        }
    }

    return res
}
