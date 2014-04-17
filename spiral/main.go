package main

import (
    "os"
    "runtime"
    "strconv"
    "math"

    "log"
    "fmt"
)

func main() {
    args := os.Args
    
    if (len(args) != 2) {
        log.Println("Usage error: /.main <some number>. Defaulting number to 24.")
        args = append(args, "24")
    }

    spiralNum, err := strconv.Atoi(args[1])
    checkError(err)
    closestSquare := findClosestSquare(spiralNum)
    x := int(math.Ceil(float64(closestSquare/2)))
    y := x
    direction := 0
    stepSize := 1
    stepThreshhold := 1

    // Just setting up our ZxZ grid. Golang doesn't allow dynamic array size allocation, so
    // we have to use slices + make (which can only allocate one level deep at a time)
    var grid = make([][]int, closestSquare)
    for i := range grid {
        grid[i] = make([]int, closestSquare)
    }

    for i := 1; i <= spiralNum; i++ {
        if (direction % 4 == 0) {
            // right
            y++
        } else if (direction % 4 == 1) {
            // down
            x++
        } else if (direction % 4 == 2) {
            //left
            y--
        } else if (direction % 4 == 3) {
            // up
            x--
        }

        grid[x][y] = i

        // Each time we've stepped <stepSize> times, change direction
        if (i % stepSize == 0) {
            direction++
        }
        
        // When we reach the next threshhold, we increase our step size and 
        // set our next threshhold (which is 2x<stepSize>, since our area
        // increases twice per 'revolution' of the square)
        if (i == stepThreshhold) {
            stepSize++
            stepThreshhold += stepSize*2
        }
    }
    
    emitGridGraphically(grid)
}

func emitGridGraphically(grid [][]int) {
    // Emit the grid as a grid (instead of a slice of slices)
    for _,innerSlice := range grid {
        for _,val := range innerSlice {
            fmt.Printf("%6d ", val);
        }
        fmt.Print("\n")
    }
}

func checkError(err error) {
    if err != nil {
        var stack [4096]byte
        runtime.Stack(stack[:], false)
        log.Printf("%q\n%s\n", err, stack[:])
    }
}

func findClosestSquare(someNum int) (int) {
    var someNumSquared float64
    var i int

    for i = someNum; someNumSquared == 0 || float64(int64(someNumSquared)) != someNumSquared && i < 100; i++ {
        someNumSquared = math.Sqrt(float64(i))
    }

    // We know that grids where our square num is even are too small, so we
    // increment those to the larger odd version
    if (int(someNumSquared) % 2 == 0) {
        return int(someNumSquared)+1
    } else {
        return int(someNumSquared)
    }
}