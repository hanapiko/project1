package main

import (
    "testing"
)

func TestCapital(t *testing.T) {
    words := []string{"hello", "world", "(cap,", "2)"}
    expected := []string{"Hello", "World"}
    result := Capital(words)
    if !equal(result, expected) {
        t.Errorf("Expecting: %v, Got: %v", expected, result)
    }
}

func TestUpp(t *testing.T) {
    words := []string{"hello", "world", "(up)"}
    expected := []string{"hello", "WORLD"}
    result := upp(words)
    if !equal(result, expected) {
        t.Errorf("Expecting: %v, Got: %v", expected, result)
    }
}

func equal(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
// import (
//     "testing"
// )

// func TestAll(t *testing.T) {
//     testCases := []struct {
//         name     string
//         expected []string
//         fn       func([]string) []string
//     }{
//         {"TestCapital", []string{"Hello", "World"}, Capital},
//         {"TestUpp", []string{"HELLO", "WORLD"}, upp},
//         {"TestLow", []string{"hello", "WORLD"}, low},
//         {"TestHex", []string{"16"}, hex},
//         {"TestBin", []string{"10"}, bin},
//         {"TestVowel", []string{"I", "am", "an", "apple"}, vowel},
//         {"TestPncs", []string{"Hello, world!"}, Pncs},
//     }

//     for _, tc := range testCases {
//         t.Run(tc.name, func(t *testing.T) {
//             words := []string{"Hello", "cap,", "2", "world"}
//             result := tc.fn(words)
//             if !equal(result, tc.expected) {
//                 t.Errorf("Expected: %v, Got: %v", tc.expected, result)
//             }
//         })
//     }
// }

// func equal(a, b []string) bool {
//     if len(a) != len(b) {
//         return false
//     }
//     for i := range a {
//         if a[i] != b[i] {
//             return false
//         }
//     }
//     return true
// }