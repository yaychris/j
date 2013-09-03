package main

import (
    "flag"
    "log"
    "os"
    "./j"
)

var (
    file  string
    dump  bool
    match string
)

func init() {
    flag.StringVar(&file, "file", "", "file name of the data set")
    flag.BoolVar(&dump, "dump", false, "dump the data file to stdout")
    flag.StringVar(&match, "match", "", "find first match")
}

func main() {
    flag.Parse()

    if file == "" {
        j.Usage()
        os.Exit(0)
    }

    set, err := j.NewJSetFromFile(file)

    if err != nil {
        log.Fatal(err)
        os.Exit(1);
    }

    if dump {
        j.Dump(set)
    } else if match != "" {
        j.Match(set, match)
    } else {
        j.Usage()
    }
}
