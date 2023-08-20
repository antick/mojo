package main

import (
	"fmt"
	"mojo/config"
)

func main() {
	fmt.Printf("Version: %s", config.Version)
	fmt.Printf("Name: %s", config.Name)
}
