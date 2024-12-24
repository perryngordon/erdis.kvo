package main 

import (
	"fmt"
	"strings"
)


// https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string
func intArrayToString(a *[]int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(*a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
