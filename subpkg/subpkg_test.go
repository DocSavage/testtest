package subpkg

import (
    "fmt"
    "testing"
)

func init() {
    ThingToSay = "Something"
}

func StuffToDo() {
    fmt.Printf("I've got stuff to do!")
}

func TestThingToSay(t *testing.T) {
    t.Logf("ThingToSay = %s\n", ThingToSay)
    if ThingToSay != "Something" {
        t.Errorf("Couldn't set package variable ThingToSay within _test.go\n")
    }
}