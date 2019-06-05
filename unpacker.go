package main


import (
    "fmt"
    "strconv"
    "strings"
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


func unpack(s string) string {
    runesList := []rune(s)
    var builder strings.Builder
    for i, r := range runesList {
        character := string(r)
        number, err := strconv.Atoi(character)
        if err == nil {
            if i == 0 {
                return "Incorrect input"
            }
            prevCharacter := string(runesList[i-1])
            if isNumber(prevCharacter) {
                return "Incorrect input"
            }
            builder.WriteString(repeat(prevCharacter, number-1))
        } else {
            builder.WriteString(character)
        }
    }
    return builder.String()
}


func main() {
    s := "45"
    res := unpack(s)
    fmt.Println(res)
}
