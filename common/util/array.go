package util

import (
    "fmt"
    "reflect"
    "strings"
)

func InArray(obj interface{}, target interface{}) bool {
    if target == nil {
        return false
    }

    targetValue := reflect.ValueOf(target)
    switch reflect.TypeOf(target).Kind() {
    case reflect.Slice, reflect.Array:
        for i := 0; i < targetValue.Len(); i++ {
            if targetValue.Index(i).Interface() == obj {
                return true
            }
        }
    case reflect.Map:
        if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
            return true
        }
    }
    return false
}

func ArrayUnique(a interface{}) (ret []interface{}) {
    ret = make([]interface{}, 0)
    va := reflect.ValueOf(a)
    for i := 0; i < va.Len(); i++ {
        v := va.Index(i).Interface()
        if !InArray(v, ret) {
            ret = append(ret, v)
        }
    }
    return ret
}

// StrArrayUnique get unique string array
func StrArrayUnique(a []string) (ret []string) {
    ret = make([]string, 0)
    length := len(a)
    for i := 0; i < length; i++ {
        if !Contains(ret, a[i]) {
            ret = append(ret, a[i])
        }
    }
    return ret
}

func IntArrayUnique(a []int64) (ret []int64) {
    unique := make(map[int64]struct{})
    for _, val := range a {
        unique[val] = struct{}{}
    }
    ret = make([]int64, len(unique))
    idx := 0
    for k := range unique {
        ret[idx] = k
        idx += 1
    }

    return ret
}

func BoolArrayUnique(a []bool) (ret []bool) {
    ret = make([]bool, 0)
    trueExist := false
    falseExist := false
    for _, item := range a {
        if item == true {
            trueExist = true
        }
        if item == false {
            falseExist = true
        }
    }
    if trueExist {
        ret = append(ret, true)
    }
    if falseExist {
        ret = append(ret, false)
    }
    return ret
}

func RemoveDuplicatesAndEmpty(slice []string) (ret []string) {
    ret = make([]string, 0)
    for _, a := range slice {
        if strings.TrimSpace(a) != "" && !Contains(ret, a) {
            ret = append(ret, a)
        }
    }
    return
}

func StrArrDiff(slice1 []string, slice2 []string) []string {
    diffStr := make([]string, 0)
    for _, i := range slice1 {
        isIn := false
        for _, j := range slice2 {
            if i == j {
                isIn = true
                break
            }
        }
        if !isIn {
            diffStr = append(diffStr, i)
        }
    }
    return diffStr
}

func IntArrIntersection(slice1 []int64, slice2 []int64) []int64 {
    intersectInt := make([]int64, 0)
    intMap := make(map[int64]bool)
    for _, i := range slice1 {
        intMap[i] = true
    }
    for _, j := range slice2 {
        if _, ok := intMap[j]; ok == true {
            intersectInt = append(intersectInt, j)
        }
    }
    return intersectInt
}

func PrettyIPStr(ips []string) string {
    if len(ips) > 2 {
        return fmt.Sprintf("%s ...", strings.Join(ips[:2], ","))
    }
    return strings.Join(ips, ",")
}

// ReverseArrayString reverse the slice's element from tail to head.
func ReverseArrayString(t []string) []string {
    if len(t) == 0 {
        return t
    }
    for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
        t[i], t[j] = t[j], t[i]
    }
    return t
}
