package main
import (
    "fmt"
)


func commandHelp(*config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println("help: Displays a help message")
    fmt.Println("exit: Exit the Pokedex")
    fmt.Println("map: Displays the next 20 locations")
    fmt.Println("mapb: Displays the previous 20 locations")
    return nil
}