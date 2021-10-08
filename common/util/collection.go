package util

import "strings"

func CalSliceDiff(oldSlice, newSlice []string) (subs, plugs []string) {
    subs = make([]string, 0)
    plugs = make([]string, 0)
    for _, a := range oldSlice {
        if !Contains(newSlice, a) {
            subs = append(subs, a)
        }
    }
    for _, b := range newSlice {
        if !Contains(oldSlice, b) {
            plugs = append(plugs, b)
        }
    }
    return
}

func CaseInsensitiveContains(s string, substr string) bool {
    return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// Contains if string target in array
func Contains(set []string, substr string) bool {
    for _, s := range set {
        if s == substr {
            return true
        }
    }
    return false
}

// ContainsInt64 if int64 target in array
func ContainsInt64(set []int64, sub int64) bool {
    for _, s := range set {
        if s == sub {
            return true
        }
    }
    return false
}

func ContainsInt(set []int64, sub int64) bool {
    for _, s := range set {
        if s == sub {
            return true
        }
    }
    return false
}

func CalSliceInt64Diff(oldSlice, newSlice []int64) (subs, inter, plugs []int64) {
    subs = make([]int64, 0)
    inter = make([]int64, 0)
    plugs = make([]int64, 0)
    for _, a := range oldSlice {
        if !ContainsInt64(newSlice, a) {
            subs = append(subs, a)
        } else {
            inter = append(inter, a)
        }
    }
    for _, b := range newSlice {
        if !ContainsInt64(oldSlice, b) {
            plugs = append(plugs, b)
        }
    }
    return
}
