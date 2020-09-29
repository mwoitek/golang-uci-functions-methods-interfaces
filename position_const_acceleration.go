package main

import (
        "fmt"
        "strconv"
)

func GenDisplaceFn(s0 float64, v0 float64, a float64) func(float64) float64 {
        DisplaceFn := func(t float64) float64 {
                              s := s0
                              s += v0 * t
                              s += 0.5 * a * t * t
                              return s
                      }
        return DisplaceFn
}

func main() {

        var s0 float64
        var v0 float64
        var a float64
        var t float64
        var err error
        var temp string

        fmt.Println("Computing the position of a particle with " +
                    "constant acceleration")

        for {
                fmt.Print("Enter the value of the initial position (in m): ")
                fmt.Scanln(&temp)
                s0, err = strconv.ParseFloat(temp, 64)
                if err == nil {
                        break
                } else {
                        fmt.Println("This isn't a valid input! Try again.")
                }
        }

        for {
                fmt.Print("Enter the value of the initial velocity (in m/s): ")
                fmt.Scanln(&temp)
                v0, err = strconv.ParseFloat(temp, 64)
                if err == nil {
                        break
                } else {
                        fmt.Println("This isn't a valid input! Try again.")
                }
        }

        for {
                fmt.Print("Enter the value of the acceleration [in m/(s^2)]: ")
                fmt.Scanln(&temp)
                a, err = strconv.ParseFloat(temp, 64)
                if err == nil {
                        break
                } else {
                        fmt.Println("This isn't a valid input! Try again.")
                }
        }

        DisplaceFn := GenDisplaceFn(s0, v0, a)
        for {
                fmt.Print("Enter the value of time (in s): ")
                fmt.Scanln(&temp)
                t, err = strconv.ParseFloat(temp, 64)
                if err == nil {
                        break
                } else {
                        fmt.Println("This isn't a valid input! Try again.")
                }
        }
        s := DisplaceFn(t)
        fmt.Printf("At t = %.2f s, the position of the particle is " +
                   "s = %.2f m.\n", t, s)

}
