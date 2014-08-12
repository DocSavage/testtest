package subpkg

import "fmt"

// If this Exported variable were in subpkg_test.go, there would be a build error even with "go test ./..."
var ThingToSay string

func SaySomethingElse() {
    fmt.Printf("In subpkg, I'd like to say: %s\n", ThingToSay)
}

