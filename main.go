package main

import (
    "fmt"
    
    "github.com/DocSavage/testtest/subpkg"
)

func MainSay() string {
    return "Something"
}

// If you run this via "go build", ThingToSay is empty, as expected since the init() 
// in subpkg_test.go isn't run.  ThingToSay is still empty when you run "go test -v ./..."
func main() {
    fmt.Printf("  Main message: %s\n", MainSay())
    fmt.Printf("Subpkg message: %s\n", subpkg.ThingToSay)
}