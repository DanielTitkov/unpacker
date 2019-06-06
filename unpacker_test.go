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
        {"abc9177d", "ab" + repeat("c", 9177) + "d"},
        {`qwe\4\5`, "qwe45"},
        {`qwe\45`, "qwe44444"},
        {`qwe\6765ee`, "qwe" + repeat("6", 765) + "ee"},
    }

    for _, tc := range testcases {
        if res := unpack(tc.inp); res != tc.outp {
            t.Errorf("Expected %s, got %s", tc.outp, res)
        }
    }
}
