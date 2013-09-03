package main

import (
    "fmt"
    "log"
    "os"
    "./j"
)

func main() {
    var file = ".j"

    set, err := j.NewJSet(file)

    if err != nil {
        log.Fatal(err)
        os.Exit(1);
    }

    for _, entry := range set.Entries {
        fmt.Println(entry)
    }
}
