package algorithm

import (
    "fmt"
    "math"
)

const (
    stars = "****************************************************************"
)

func printMatrix(matrix [][]int) {
    fmt.Println(stars)
    fmt.Println()
    for _, row := range matrix {
        for _, val := range row {
            fmt.Printf("%v\t", val)
        }
        fmt.Println()
    }
    fmt.Println()
    fmt.Println(stars)
}

// Solve - solves assignment problem with Hungarian method
func Solve(baseMatrix [][]int, maximize bool) {
    matrix := make([][]int, len(baseMatrix))
    for i := range matrix {
        matrix[i] = make([]int, len(baseMatrix[i]))
    }

    for i := range matrix {
        for j := range matrix[i] {
            matrix[i][j] = baseMatrix[i][j]
        }
    }
    if maximize {
        max := 0
        for _, row := range matrix {
            for _, val := range row {
                if max < val {
                    max = val
                }
            }
        }
        for _, row := range matrix {
            for j := range row {
                row[j] = int(math.Abs(float64(row[j] - max)))
            }
        }
    }
    fmt.Println("Base matrix:")
    if maximize {
        printMatrix(baseMatrix)
        fmt.Println("Converted to minimize:")
        printMatrix(matrix)
    } else {
        printMatrix(matrix)
    }

    fmt.Println("Row Reduction:")
    matrix = reduceRow(matrix)
    printMatrix(matrix)


    fmt.Println("Column Reduction:")
    matrix = reduceColumn(matrix)
    printMatrix(matrix)

    coordinates := scanRow(matrix)

    if !allZerosAreDestroyed(matrix, coordinates) {
        coordinates = scanColumn(matrix, coordinates)
    }
    fmt.Println(coordinates)

    if len(coordinates) != len(matrix) {
        for {
            aliveMin := findMinFromAliveCells(matrix, coordinates)
            fmt.Println("aliveMin: ", aliveMin)
            addAliveMinToIntersections(matrix, coordinates, aliveMin)
            subtractFromAlives(matrix, coordinates, aliveMin)
            printMatrix(matrix)

            coordinates = scanRow(matrix)
            if !allZerosAreDestroyed(matrix, coordinates) {
                coordinates = scanColumn(matrix, coordinates)
            }
            printMatrix(matrix)
            fmt.Println(coordinates)
            if len(coordinates) == len(matrix) {
                break
            }
        }
    }

    coordinates = sortCoordinates(coordinates)

    fmt.Printf("\nJob\t\tOperator\tCost\n")
    for _, coord := range coordinates {
        fmt.Printf("%v\t\t%v\t\t%v\n", coord.row, coord.column,
            baseMatrix[coord.row][coord.column])
    }
    printMatrix(baseMatrix)
    fmt.Println(coordinates)
}

func sortCoordinates(coordinates []coordinate) []coordinate {
    for i := range coordinates {
        for j := range coordinates {
            if coordinates[i].row < coordinates[j].row {
                temp := coordinates[i]
                coordinates[i] = coordinates[j]
                coordinates[j] = temp
            }
        }
    }

    return coordinates
}
// reduceRow - Step 1: Row and Column Reduction(subtract minimum value
// of the row from the entries of that row)
func reduceRow(matrix [][]int) [][]int {
    minimals := make([]int, 0)
    for _, row := range matrix {
        min := row[0]
        for _, val := range row {
            if val < min {
                min = val
            }
        }
        minimals = append(minimals, min)
    }

    for i, row := range matrix {
        for j := range row {
            matrix[i][j] -= minimals[i]
        }
    }

    return matrix
}

// reduceColumn - Step 1: Row and Column Reduction(subtract minimum value
// of the column from the entries of that column)
func reduceColumn(matrix [][]int) [][]int {
    minimals := make([]int, 0)
    for i := 0; i < len(matrix[0]); i++ {
        min := 1000000000
        for j := 0; j < len(matrix); j++ {
            if matrix[j][i] < min {
                min = matrix[j][i]
            }
        }
        minimals = append(minimals, min)
    }

    for i := 0; i < len(matrix[0]); i++ {
        for j := 0; j < len(matrix); j++ {
            matrix[j][i] -= minimals[i]
        }
    }

    return matrix
}

type coordinate struct {
    row, column int
    vertical bool
}

func scanRow(matrix [][]int) []coordinate {
    coordinates := make([]coordinate, 0)

    for i, row := range matrix {
        zeroCount := 0
        columnCoord := 0
        for j := range row {
            canProceed := true
            for _, coord := range coordinates {
                if j == coord.column {
                    canProceed = false
                }
            }
            if !canProceed {
                continue
            }
            if row[j] == 0 {
                zeroCount++
                columnCoord = j
            }
        }

        if zeroCount == 1 {
            coordinates = append(coordinates, coordinate{i, columnCoord, true})
        }
    }

    return coordinates
}

func scanColumn(matrix [][]int, coordinates []coordinate) []coordinate {
    for i := 0; i < len(matrix[0]); i++ {
        canProceed := true
        for _, coord := range coordinates {
            if coord.column == i {
                canProceed = false
            }
        }
        if !canProceed {
            continue
        }

        zeroCount := 0
        rowCoord := 0
        for j := 0; j < len(matrix); j++ {
            for _, coord := range coordinates {
                if coord.row == j && !coord.vertical {
                    canProceed = false
                }
            }
            if !canProceed {
                continue
            }

            if matrix[j][i] == 0 {
                zeroCount++
                rowCoord = j
            }
        }

        if zeroCount == 1 {
            coordinates = append(coordinates, coordinate{rowCoord, i, false})
        }
    }

    return coordinates
}

func allZerosAreDestroyed(matrix [][]int, coordinates []coordinate) bool {
    for _, row := range matrix {
        for j, val := range row {
            canProceed := true
            for _, coord := range coordinates {
                if coord.column == j {
                    canProceed = false
                }
            }
            if !canProceed {
                continue
            }

            if val == 0 {
                return false
            }
        }
    }

    return true
}

func findMinFromAliveCells(matrix [][]int, coordinates []coordinate) int {
    //if len(coordinates) == 0 {
      //  return 0;
    //}
    min := 1000000000
    for i, row := range matrix {
        canProceed := true
        for _, coord := range coordinates {
            if coord.row == i && !coord.vertical {
                canProceed = false
                break
            }
        }
        if !canProceed {
            continue
        }

        for j, val := range row {
            for _, coord := range coordinates {
                if coord.column != j {
                    canProceed = true
                } else if coord.column == j && coord.vertical {
                    canProceed = false
                    break
                }
            }
            if !canProceed {
                continue
            }

            if min > val {
                min = val
            }
        }
    }

    return min
}

func addAliveMinToIntersections(matrix [][]int,
    coordinates []coordinate, aliveMin int) {
    for _, coord := range coordinates {
        if !coord.vertical {
            for _, coord2 := range coordinates {
                if coord2.vertical {
                    matrix[coord.row][coord2.column] += aliveMin
                }
            }
        }
    }
}

func subtractFromAlives(matrix [][]int,
    coordinates []coordinate, aliveMin int) {

    for i, row := range matrix {
        canProceed := true
        for _, coord := range coordinates {
            if coord.row == i && !coord.vertical {
                canProceed = false
                break
            }
        }
        if !canProceed {
            continue
        }

        for j := range row {
            for _, coord := range coordinates {
                if coord.column != j {
                    canProceed = true
                } else if coord.column == j && coord.vertical {
                    canProceed = false
                    break
                }
            }
            if !canProceed {
                continue
            }

            row[j] -= aliveMin
        }
    }
}
