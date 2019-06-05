package main


import "testing"


func TestRepeat(t *testing.T) {
    t.Log("Repeating the string n times")

    type testcase struct {
        inp string
        number int
        outp string
    }

    testcases := []testcase {
        {"foo", 0, ""},
        {"foo", 1, "foo"},
        {"foo", 3, "foofoofoo"},
    }

    for _, tc := range testcases {
        if res := repeat(tc.inp, tc.number); res != tc.outp {
            t.Errorf("Expected %s, got %s", tc.outp, res)
        }
    }
}


func TestSimpleUnpacks(t *testing.T) {
    t.Log("Trying string with single numbers without escaping")

    type testcase struct {
        inp string
        outp string
    }

    testcases := []testcase {
        {"a4bc2d5e", "aaaabccddddde"},
        {"abcd", "abcd"},
        {"45", "Incorrect input"},
    }

    for _, tc := range testcases {
        if res := unpack(tc.inp); res != tc.outp {
            t.Errorf("Expected %s, got %s", tc.outp, res)
        }
    }
}
