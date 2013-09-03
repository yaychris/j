package j

import (
  "io/ioutil"
  "strconv"
  "strings"
)

type JSet struct {
    Entries []*JEntry
}

func NewJSet(file string) (*JSet, error) {
    rawBytes, err := ioutil.ReadFile(file)

    if err != nil {
        return nil, err
    }

    text := string(rawBytes)

    lines := strings.Split(text, "\n")

    set := &JSet{ Entries: make([]*JEntry, 0, len(lines)) }

    for _, line := range lines {
        fields := strings.Split(line, "|")

        if len(fields) == 3 {
            path         := fields[0]
            rank, _      := strconv.ParseFloat(fields[1], 64)
            timestamp, _ := strconv.Atoi(fields[2])

            entry := NewJEntry(path, rank, timestamp)

            set.Entries = append(set.Entries, entry)
        }
    }

    return set, nil
}
