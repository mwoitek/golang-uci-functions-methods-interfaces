package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
)

// Definition of the "Animal" interface.
type Animal interface {
    Eat()
    Move()
    Speak()
}

// "Cow" type and methods associated with it.
type Cow struct {}

func (c Cow) Eat() {
    fmt.Printf("This animal eats grass.\n\n")
}

func (c Cow) Move() {
    fmt.Printf("This animal moves by walking.\n\n")
}

func (c Cow) Speak() {
    fmt.Printf("The sound this animal makes is moo.\n\n")
}

// "Bird" type and methods associated with it.
type Bird struct {}

func (b Bird) Eat() {
    fmt.Printf("This animal eats worms.\n\n")
}

func (b Bird) Move() {
    fmt.Printf("This animal moves by flying.\n\n")
}

func (b Bird) Speak() {
    fmt.Printf("The sound this animal makes is peep.\n\n")
}

// "Snake" type and methods associated with it.
type Snake struct {}

func (s Snake) Eat() {
    fmt.Printf("This animal eats mice.\n\n")
}

func (s Snake) Move() {
    fmt.Printf("This animal moves by slithering.\n\n")
}

func (s Snake) Speak() {
    fmt.Printf("The sound this animal makes is hsss.\n\n")
}

// This function creates an animal.
func create(opt string) Animal {

    var a Animal

    fmt.Printf("Created it!\n\n")

    switch opt {
        case "cow":
            a = Cow{}
        case "bird":
            a = Bird{}
        case "snake":
            a = Snake{}
    }

    return a

}

// This function retrieves the information the user wants.
func retrieve(a Animal, opt string) {
    switch opt {
        case "eat":
            a.Eat()
        case "move":
            a.Move()
        case "speak":
            a.Speak()
    }
}

// This function prints the initial message to the user.
func init_msg() {
    fmt.Print("Type your commands!\n")
    fmt.Print("When you're done, enter q to quit.\n\n")
}

// This function returns three maps. The corresponding keys are strings
// containing the valid options for the user.
func valid_opts() (map[string] bool, map[string] bool, map[string] bool) {

    // The keys of this map are the available commands.
    opts_c := make(map[string] bool)
    opts_c["newanimal"] = true
    opts_c["query"] = true

    // The keys of this map are the types of animal the user can create.
    opts_a := make(map[string] bool)
    opts_a["cow"] = true
    opts_a["bird"] = true
    opts_a["snake"] = true

    // The keys of this map are the types of information the user can retrieve.
    opts_i := make(map[string] bool)
    opts_i["eat"] = true
    opts_i["move"] = true
    opts_i["speak"] = true

    return opts_c, opts_a, opts_i

}

// This function reads, processes and validates the user's input.
func read(opts_c, opts_a, opts_i map[string] bool,
          animals map[string] Animal) (string, string, string, error) {

    var key_exists bool

    // Reads and processes the user's input.
    fmt.Print("> ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")
    input = strings.Trim(input, " ")
    input = strings.ToLower(input)

    // Checks if the user wants to quit.
    if strings.Compare(input, "q") == 0 {
        err := errors.New("Quitting.\n")
        return "", "", "", err
    }

    // Validates the user's input.

    input_pts := strings.Fields(input)

    if len(input_pts) < 3 {
        err := errors.New("Not enough arguments!\n\n")
        return "", "", "", err
    }

    optc := input_pts[0]
    _, key_exists = opts_c[optc]
    if !key_exists {
        err := errors.New("Command unavailable!\n\n")
        return "", "", "", err
    }

    name := input_pts[1]
    _, key_exists = animals[name]
    if strings.Compare(optc, "newanimal") == 0 {
        if key_exists {
            err := errors.New("An animal with this name already exists!\n\n")
            return "", "", "", err
        }
    } else {
        if !key_exists {
            err := errors.New("There's no animal with this name!\n\n")
            return "", "", "", err
        }
    }

    opt := input_pts[2]
    if strings.Compare(optc, "newanimal") == 0 {
        _, key_exists = opts_a[opt]
        if !key_exists {
            err := errors.New("Animal unavailable!\n\n")
            return "", "", "", err
        }
    } else {
        _, key_exists = opts_i[opt]
        if !key_exists {
            err := errors.New("Information unavailable!\n\n")
            return "", "", "", err
        }
    }

    return optc, name, opt, nil

}

func main() {

    var optc string
    var name string
    var opt string
    var err error

    // Prints the initial message to the user.
    init_msg()

    // Map for storing the user-defined animals.
    animals := make(map[string] Animal)

    // Creates three maps containing the valid options for the user.
    opts_c, opts_a, opts_i := valid_opts()

    // Interacts with the user.
    for {
        optc, name, opt, err = read(opts_c, opts_a, opts_i, animals)
        if err != nil {
            fmt.Print(err)
            if strings.Compare(err.Error(), "Quitting.\n") == 0 {
                break
            }
        } else {
            if strings.Compare(optc, "newanimal") == 0 {
                animals[name] = create(opt)
            } else {
                retrieve(animals[name], opt)
            }
        }
    }

}
