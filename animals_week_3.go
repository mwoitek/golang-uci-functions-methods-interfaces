package main

import (
        "bufio"
        "errors"
        "fmt"
        "os"
        "strings"
)

// Definition of the "Animal" type.
type Animal struct {
        food string
        locomotion string
        noise string
}

// This method prints the food a given animal eats.
func (a Animal) Eat() {
        fmt.Printf("This animal eats %s.\n\n", a.food)
}

// This method prints a given animal's means of locomotion.
func (a Animal) Move() {
        fmt.Printf("This animal moves by %sing.\n\n", a.locomotion)
}

// This method prints the sound a given animal makes.
func (a Animal) Speak() {
        fmt.Printf("The sound this animal makes is %s.\n\n", a.noise)
}

// This method calls the method specified by the user.
func (a Animal) call_specified_method(opti string) {
        switch opti {
                case "eat":
                        a.Eat()
                case "move":
                        a.Move()
                case "speak":
                        a.Speak()
        }
}

// This function returns an object of the "Animal" type containing the
// information about cows.
func info_cow() Animal {
        var cow Animal
        cow.food = "grass"
        cow.locomotion = "walk"
        cow.noise = "moo"
        return cow
}

// This function returns an object of the "Animal" type containing the
// information about birds.
func info_bird() Animal {
        var bird Animal
        bird.food = "worms"
        bird.locomotion = "fly"
        bird.noise = "peep"
        return bird
}

// This function returns an object of the "Animal" type containing the
// information about snakes.
func info_snake() Animal {
        var snake Animal
        snake.food = "mice"
        snake.locomotion = "slither"
        snake.noise = "hsss"
        return snake
}

// This function returns a map. This map stores three objects of the "Animal"
// type. These objects contain the information about cows, birds and snakes.
func info_animals() map[string] Animal {
        animals := make(map[string] Animal)
        animals["cow"] = info_cow()
        animals["bird"] = info_bird()
        animals["snake"] = info_snake()
        return animals
}

// This function prints the initial message to the user.
func init_msg() {
        fmt.Print("Welcome!\n\n")
        fmt.Print("You have 2 options:\n\n")
        fmt.Println("1 - Enter the name of an animal and " +
                    "the desired information about this animal.")
        fmt.Println("The options for the name of the animal are")
        fmt.Println("a - cow")
        fmt.Println("b - bird")
        fmt.Println("c - snake")
        fmt.Println("The options for retrieving information about an animal are")
        fmt.Println("i - eat")
        fmt.Println("ii - move")
        fmt.Println("iii - speak")
        fmt.Println("For instance, to know what a bird eats, you need to type")
        fmt.Print("bird eat\n\n")
        fmt.Print("2 - Type q to quit.\n\n")
}

// This function returns two maps whose keys are strings containing the valid
// options for the user.
func gen_valid_opts() (map[string] bool, map[string] bool) {
        opts_animal := make(map[string] bool)
        opts_animal["cow"] = true
        opts_animal["bird"] = true
        opts_animal["snake"] = true
        opts_info := make(map[string] bool)
        opts_info["eat"] = true
        opts_info["move"] = true
        opts_info["speak"] = true
        return opts_animal, opts_info
}

// This function reads and processes the user's input. Two of the return values
// are strings that store the user's input after processing. The function also
// returns an error variable. For convenience, this function handles the user's
// decision to quit the program as an error.
func read_input(opts_animal, opts_info map[string] bool) (string, string, error) {

        var valid bool
        fmt.Print("> ")

        // Reads and processes the user's input.
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSuffix(input, "\n")
        input = strings.Trim(input, " ")
        input = strings.ToLower(input)

        // Checks if the user wants to quit.
        if strings.Compare(input, "q") == 0 {
                err := errors.New("Quitting.\n")
                return "", "", err
        }

        // Validates the user's input.
        input_pts := strings.Fields(input)
        if len(input_pts) < 2 {
                err := errors.New("To retrieve information about an animal, " +
                                  "you need to specify two options! Try again.\n\n")
                return "", "", err
        }
        opta := input_pts[0]
        _, valid = opts_animal[opta]
        if !valid {
                err := errors.New("Invalid animal name! Try again.\n\n")
                return "", "", err
        }
        opti := input_pts[1]
        _, valid = opts_info[opti]
        if !valid {
                err := errors.New("The requested information isn't available! " +
                                  "Try again.\n\n")
                return "", "", err
        }

        return opta, opti, nil

}

func main() {

        var err error
        var opta string
        var opti string

        // Prints the initial message to the user.
        init_msg()

        // Creates a map containing the information about the animals.
        animals := info_animals()

        // Creates two maps containing the valid options for the user.
        opts_animal, opts_info := gen_valid_opts()

        // Interacts with the user.
        for {
                opta, opti, err = read_input(opts_animal, opts_info)
                if err != nil {
                        fmt.Print(err)
                        if strings.Compare(err.Error(), "Quitting.\n") == 0 {
                                break
                        }
                } else {
                        animals[opta].call_specified_method(opti)
                }
        }

}
