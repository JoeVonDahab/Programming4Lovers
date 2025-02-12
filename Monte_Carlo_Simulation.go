package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Compute the house edge for Craps
func ComputeCrapsHouseEdge(numTrials int) float64 {
    wins := 0

    for i := 0; i < numTrials; i++ {
        if PlayCrapsOnce() {
            wins++
        }
    }

    losses := numTrials - wins
    return float64(losses) / float64(numTrials) // House edge formula
}

// Simulates one game of Craps
func PlayCrapsOnce() bool {
    firstRoll := SumDice(2)

    if firstRoll == 7 || firstRoll == 11 {
        return true
    } else if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
        return false
    }

    for {
        newRoll := SumDice(2)
        if newRoll == firstRoll {
            return true
        } else if newRoll == 7 {
            return false
        }
    }
}

// Rolls multiple dice and returns the sum
func SumDice(numDice int) int {
    sum := 0
    for i := 0; i < numDice; i++ {
        sum += RollDie()
    }
    return sum
}

// Simulates rolling a six-sided die
func RollDie() int {
    return rand.Intn(6) + 1
}

func main() {
    rand.Seed(time.Now().UnixNano()) // Ensures different results on each run
    numTrials := 1000000 // Large number of trials for accuracy
    houseEdge := ComputeCrapsHouseEdge(numTrials)
    fmt.Printf("House Edge: %.5f\n", houseEdge)
}
