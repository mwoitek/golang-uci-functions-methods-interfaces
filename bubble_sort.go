package main

import (
        "fmt"
        "strconv"
        "strings"
)

func read_nums() (*[]int, int) {

        var err error
        var input_str string
        L := 0
        var new_int int
        nums := make([]int, 0)

        fmt.Println("Begin by specifying a set of integers.")
        fmt.Println("You can enter up to 10 numbers.")
        for L < 10 {
                fmt.Print("Enter an integer or X to quit: ")
                fmt.Scanf("%s", &input_str)
                if strings.Compare(strings.ToLower(input_str), "x") == 0 {
                        break
                } else {
                        new_int, err = strconv.Atoi(input_str)
                        if err != nil {
                                fmt.Println("Invalid input!")
                        } else {
                                L++
                                nums = append(nums, new_int)
                        }
                }
        }

        return &nums, L

}

func Swap(nums *[]int, i int) {

        i1 := i + 1
        temp := (*nums)[i]
        (*nums)[i] = (*nums)[i1]
        (*nums)[i1] = temp

}

func BubbleSort(nums *[]int, L int) {

        var i1 int
        var new_L int

        for {
                new_L = 0
                for i := 1; i < L; i++ {
                        i1 = i - 1
                        if (*nums)[i1] > (*nums)[i] {
                                Swap(nums, i1)
                                new_L = i
                        }
                }
                L = new_L
                if L <= 1 {
                        break
                }
        }

}

func main() {

        n, L := read_nums()
        if L == 0 {
                fmt.Println("Your list of integers is empty.")
                fmt.Println("There is nothing to sort.")
        } else {
                BubbleSort(n, L)
                fmt.Println("Sorted list:")
                fmt.Println(*n)
        }

}
