package main


import "testing"


func TestRepeat(t *testing.T) {
    t.Log("Repeating the string n times")

    type testcaseRep struct {
        number int
        inp string
        outp string
    }

    testcases := []testcaseRep {
        {0, "foo", ""},
        {1, "foo", "foo"},
        {3, "foo", "foofoofoo"},
    }

    for _, tc := range testcases {
        if res := repeat(tc.inp, tc.number); res != tc.outp {
            t.Errorf("Expected %s, got %s", tc.outp, res)
        }
    }
}


func TestSimpleUnpacks(t *testing.T) {
    t.Log("Trying string without escaping")

    type testcase struct {
        inp string
        outp string
    }

    testcases := []testcase {
        {"a4bc2d5e", "aaaabccddddde"},
        {"abcd", "abcd"},
        {"45", "Incorrect input"},
        {"a10", "aaaaaaaaaa"},
        {"abc9177d", "ab" + repeat("c", 9177) + "d"},
    }

    for _, tc := range testcases {
        if res := unpack(tc.inp); res != tc.outp {
            t.Errorf("Expected %s, got %s", tc.outp, res)
        }
    }
}
