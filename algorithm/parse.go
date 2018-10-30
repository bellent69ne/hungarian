package algorithm

import (
    "strconv"
    "io/ioutil"
    "log"
    "strings"
)

func Process(filename string) [][]int {
    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    slice := strings.Split(string(contents), "\n")

    matrix := make([][]int, 0)
    numSlice := make([]int, 0)
    for _, val := range slice {
        splitted := strings.Split(val, " ")
        for _, str := range splitted {
            if str != "" {
                num, err := strconv.Atoi(str)
                if err != nil {
                    log.Fatal(err)
                }
                numSlice = append(numSlice, num)
            }
        }
        matrix = append(matrix, numSlice)
        numSlice = nil
    }

    matrix = matrix[0:len(matrix) - 1]

    return matrix
}
