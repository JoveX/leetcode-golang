package main

import "fmt"

func main() {

    s := "[])"
    res := isValid2(s);
    fmt.Println(res)
}

func isValid2(s string) bool {
    s_len := len(s)

    if s_len % 2 != 0 {
        return false
    }

    res := []string{}
    for i := 0; i < s_len; i++ {
        str := s[i:i+1]
        if str == "}" || str == "]" || str == ")" {
            res_len := len(res)
            if res_len < 1 {
                return false
            }
            l_str := ""
            switch str {
                case "}":
                    l_str = "{"
                    break
                case "]":
                    l_str = "["
                    break
                case ")":
                    l_str = "("
                    break
            }
            if res[res_len - 1] == l_str {
                res = res[:res_len - 1]
            } else {
                return false
            }
        } else {
            res = append(res, str)
        }
    }

    if len(res) > 0 {
        return false
    } else {
        return true
    }

}


func isValid1(s string) bool {
    s_len := len(s)

    if s_len % 2 != 0 {
        return false
    }

    res := ""
    for i := 0; i < s_len; i++ {
        str := s[i:i+1]
        if str == "}" || str == "]" || str == ")" {
            res_len := len(res)
            if res_len < 1 {
                return false
            }
            l_str := ""
            switch str {
                case "}":
                    l_str = "{"
                    break
                case "]":
                    l_str = "["
                    break
                case ")":
                    l_str = "("
                    break
            }
            if res[res_len - 1:res_len] == l_str {
                res = res[0:res_len - 1]
            } else {
                return false
            }
        } else {
            res = res + str
        }
    }

    if len(res) > 0 {
        return false
    } else {
        return true
    }

}
