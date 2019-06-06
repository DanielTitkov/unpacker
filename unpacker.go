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


func resolveBuffer(elem, bufs string) string {
    firstChar := bufs[0:1]
    remainder := bufs[1:]
    switch {
    case firstChar != `\`:
        charNumber := toNumber(bufs)
        return repeat(elem, charNumber)
    case len(bufs) == 2:
        return elem + remainder
    case len(bufs) > 2:
        return elem + resolveBuffer(remainder[0:1], remainder[1:])
    default:
        return elem
    }
}


func unpack(s string) string {
    runesList := []rune(s)
    var result, buffer strings.Builder
    for i := len(runesList)-1; i >= 0; i-- {
        elem := string(runesList[i])
        switch {
        case isNumber(elem) || elem == `\`:
            buffer.WriteString(elem)
        case buffer.Len() > 0:
            bufferString := resolveBuffer(elem, strrev.Reverse(buffer.String()))
            result.WriteString(strrev.Reverse(bufferString))
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
    s := "a4bc2d5e"
    res := unpack(s)
    fmt.Println(s, "=>", res)
}
