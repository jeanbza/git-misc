package main

import (
    "os"
    "runtime"
    "strconv"
    "math"

    "log"
    "fmt"
)

// SOLUTION OUTLINE
// Our basic approach is to go one direction a certain step size, change direction
// and go the same step size, then increase our step size and repeat. The direction 
// cycles between right, down, left, up each <stepsize> steps
// 
// SOLUTION EFFICIENCY
// This solution runs with O(n) space and time complexity
func main() {
    args := os.Args
    
    if (len(args) != 2) {
        log.Println("Usage error: /.main <some number>. Defaulting number to 24.")
        args = append(args, "24")
    }

    spiralNum, err := strconv.Atoi(args[1])
    checkError(err)
    closestSquare := findClosestSquare(spiralNum)
    x := int(math.Floor(closestSquare/2))
    y := x
    direction := 0
    stepSize := 1
    stepThreshhold := 1

    // Setting up our ZxZ grid. Golang doesn't allow dynamic array size allocation, so
    // we have to use slices + make (which can only allocate one level deep at a time)
    var grid = make([][]int, int(closestSquare))
    for i := range grid {
        grid[i] = make([]int, int(closestSquare))
    }

    for i := 1; i <= spiralNum; i++ {
        if (direction % 4 == 0) {
            y++ // right
        } else if (direction % 4 == 1) {
            x++ // down
        } else if (direction % 4 == 2) {
            y-- //left
        } else if (direction % 4 == 3) {
            x-- // up
        }

        grid[x][y] = i

        // Each time we've stepped <stepSize> times, change direction
        if (i % stepSize == 0) {
            direction++
        }
        
        // At the next threshhold, we increase our step size and set our next threshhold
        if (i == stepThreshhold) {
            stepSize++
            stepThreshhold += stepSize*2
        }
    }
    
    emitGridGraphically(grid)
}

// Emits slice of slices as a grid
func emitGridGraphically(grid [][]int) {
    for _, innerSlice := range grid {
        for _, val := range innerSlice {
            fmt.Printf("%6d ", val);
        }
        fmt.Print("\n")
    }
}

// Finds the closest square (moving up only) to someNum
func findClosestSquare(someNum int) (float64) {
    // In the case of 1, we want a 3x3 grid. For all other numbers the following logic will work
    if (someNum == 1) {
        return 3
    }

    var someNumSquared float64
    var i int

    for i = someNum; someNumSquared == 0 || float64(int64(someNumSquared)) != someNumSquared && i < 100; i++ {
        someNumSquared = math.Sqrt(float64(i))
    }

    // We know that grids where our square num is even are too small, so we increment those
    // to the larger odd version
    if (int(someNumSquared) % 2 == 0) {
        return someNumSquared+1
    } else {
        return someNumSquared
    }
}

// Stacktrace functionality for error handling
func checkError(err error) {
    if err != nil {
        var stack [4096]byte
        runtime.Stack(stack[:], false)
        log.Printf("%q\n%s\n", err, stack[:])
    }
}