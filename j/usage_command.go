package j

import "fmt"

func Usage() error {
    fmt.Println("usage:")
    fmt.Println("  jgo --file=<file> --dump                Dump the data file to stdout")
    fmt.Println("  jgo --file=<file> --add=<path>          Add a new entry to the data file")
    fmt.Println("  jgo --file=<file> --complete <regex>    Find a path matching <regex>")

    return nil
}
