package main

import (
    "testing"
    
    "github.com/DocSavage/testtest/subpkg"
)

func TestSaySomething(t *testing.T) {
    if MainSay() != subpkg.ThingToSay {
        t.Errorf("Didn't say %s!\n", subpkg.ThingToSay)
    }
    
    // When running "go test -v ./...", the ThingToSay is still empty,
    // which suggests that the init() in subpkg_test.go is not setting
    // that exported variable.
    subpkg.SaySomethingElse()
    
    // If uncomment below, get a build error with "go test ./...",
    // so exported functions with subpkg_test.go aren't visible.
    
    //subpkg.StuffToDo()
}