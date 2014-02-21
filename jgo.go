package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "./j"
)

var (
    file       string
    dump       bool
    match      string
    matchLimit int
    pathToAdd  string
)

func init() {
    file = os.Getenv("J_DATA")

    flag.BoolVar(&dump,        "dump",  false, "dump the data file to stdout")
    flag.StringVar(&match,     "match", "",    "find first match")
    flag.IntVar(&matchLimit,   "limit", -1,    "limit the number of matches")
    flag.StringVar(&pathToAdd, "add",   "",    "path to add to the database")
}

func main() {
    flag.Parse()

    if file == "" {
        fmt.Println("must specify the data file in the J_DATA environment variable")
        os.Exit(1)
    }

    set, err := j.NewJSetFromFile(file)

    if err != nil {
        log.Fatal(err)
        os.Exit(1);
    }

    if dump {
        j.Dump(set)
    } else if match != "" {
        j.Match(set, match, matchLimit)
    } else if pathToAdd != "" {
        j.Add(set, pathToAdd, file)
    } else {
        j.Usage()
    }
}
