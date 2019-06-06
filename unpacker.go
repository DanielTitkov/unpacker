package main


import (
    "fmt"
    "strconv"
    "strings"
    "log"
    "4d63.com/strrev"
)


func repeat(s string, n int) string {
    var builder strings.Builder
    for i := 0; i < n; i++ {
        builder.WriteString(s)
    }
    return builder.String()
}


func isNumber(s string) bool {
    _, err := strconv.Atoi(s)
    return err == nil
}


func toNumber(s string) int {
    i, err := strconv.Atoi(s)
    if err == nil {
        return i
    }
    log.Printf("Not a number: %s", s)
    return 1
}


func unpack(s string) string {
    runesList := []rune(s)
    var result, buffer strings.Builder
    for i := len(runesList)-1; i >= 0; i-- {
        elem := string(runesList[i])
        switch {
        case isNumber(elem):
            buffer.WriteString(elem)
        case elem == `\` && buffer.Len() > 1:
            charNumber := toNumber(strrev.Reverse(buffer.String())[1:])
            result.WriteString(repeat(strrev.Reverse(buffer.String())[0:1], charNumber))
            buffer.Reset()
        case elem == `\` && buffer.Len() == 1:
            result.WriteString(buffer.String())
            buffer.Reset()
        case buffer.Len() > 0:
            charNumber := toNumber(strrev.Reverse(buffer.String()))
            result.WriteString(repeat(elem, charNumber))
            buffer.Reset()
        default:
            result.WriteString(elem)
        }
    }
    if buffer.Len() != 0 {
        return "Incorrect input"
    }
    return strrev.Reverse(result.String())
}


func main() {
    s := `a4b\\54c2d5e`
    res := unpack(s)
    fmt.Println(s)
    fmt.Println(res)
}
