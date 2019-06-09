package main


import (
    "testing"
    "strings"
)


func TestBufferResolve(t *testing.T) {
    t.Log("Resolving buffer data")

    type testcaseBuf struct {
        elem string
        bufs string
        outp string
    }

    testcases := []testcaseBuf {
        {"e", "4", "eeee"},
        {"e", "457", strings.Repeat("e", 457)},
        {"e", `\5`, "e5"},
        {"e", `\\`, `e\`},
        {"e", `\54`, "e5555"},
        {"e", `\\4`, `e\\\\`},
    }

    for _, tc := range testcases {
        if res := resolveBuffer(tc.elem, tc.bufs); res != tc.outp {
            t.Errorf("Expected %s, got %s", tc.outp, res)
        }
    }
}


func TestUnpacks(t *testing.T) {
    t.Log("Trying string unpacking")

    type testcase struct {
        inp string
        outp string
    }

    testcases := []testcase {
        {"a4bc2d5e", "aaaabccddddde"},
        {"abcd", "abcd"},
        {"45", "Incorrect input"},
        {"a10", "aaaaaaaaaa"},
        {"abc9177d", "ab" + strings.Repeat("c", 9177) + "d"},
        {`qwe\4\5`, "qwe45"},
        {`qwe\45`, "qwe44444"},
        {`qwe\6765ee`, "qwe" + strings.Repeat("6", 765) + "ee"},
        {`qwe\\5`, `qwe\\\\\`},
        {`aa\\bb`, `aa\bb`},
        {`aa\bb`, `aabb`},
    }

    for _, tc := range testcases {
        if res := unpack(tc.inp); res != tc.outp {
            t.Errorf("Expected %s, got %s", tc.outp, res)
        }
    }
}
