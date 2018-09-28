package main

import "fmt"

func main() {
    digits := "23"
    res := letterCombinations(digits)
    fmt.Println(':', res)
}


func letterCombinations(digits string) []string {
    digits_len := len(digits)

    letter_map := make(map[string][]string)
    letter_map["2"] = []string{"a", "b", "c"}
    letter_map["3"] = []string{"d", "e", "f"}
    letter_map["4"] = []string{"g", "h", "i"}
    letter_map["5"] = []string{"j", "k", "l"}
    letter_map["6"] = []string{"m", "n", "o"}
    letter_map["7"] = []string{"p", "q", "r", "s"}
    letter_map["8"] = []string{"t", "u", "v"}
    letter_map["9"] = []string{"w", "x", "y", "z"}

    res := []string{}
    num_str := ""
    for i := 0; i < digits_len; i++ {
        num_str = digits[i:i+1]

        letters := letter_map[num_str]

        if len(res) < 1 {
            res = letters
            continue;
        }

        tmp_arr := []string{}

        for j := 0; j < len(res); j++ {
            for k := 0; k < len(letters); k++ {
                tmp_arr = append(tmp_arr, res[j] + letters[k])
            }
        }

        res = tmp_arr
    }

    return res
}
