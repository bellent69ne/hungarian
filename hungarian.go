package main

import (
    "github.com/bellent69ne/hungarian/algorithm"
    "fmt"
    "os"
)

func main() {
    //matrix := [][]int{{9,  11, 14, 11, 7},
      //                {6,  15, 13, 13, 10},
        //              {12, 13, 6,  8,  8},
          //            {11, 9,  10, 12, 9},
            //          {7,  12, 14, 10, 14}}



    matrix := algorithm.Process(os.Args[1])
    switch os.Args[2] {
    case "max":
        algorithm.Solve(matrix, true)
    case "min":
        algorithm.Solve(matrix, false)
    default:
        {
            fmt.Println("Invalid arguments...")
            return
        }
    }
}
