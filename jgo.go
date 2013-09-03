package main

import (
    "flag"
    "log"
    "os"
    "./j"
)

func main() {
    file := flag.String("file", "", "file name of the data set")
    dump := flag.Bool("dump", false, "dump the data file to stdout")

    flag.Parse()

    if *file == "" {
        j.Usage()
        os.Exit(0)
    }

    set, err := j.NewJSetFromFile(*file)

    if err != nil {
        log.Fatal(err)
        os.Exit(1);
    }

    if *dump {
        j.Dump(set)
    } else {
        j.Usage()
    }
}
