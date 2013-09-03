package j

import (
    "io/ioutil"
    "strconv"
    "strings"
)

type JSet struct {
    Entries []*JEntry
}

type EachFunc   func(entry *JEntry)
type SelectFunc func(entry *JEntry) bool

func NewJSetFromFile(file string) (*JSet, error) {
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

func (set *JSet) Each(f EachFunc) {
    for i := range set.Entries {
        f(set.Entries[i])
    }
}

func (set *JSet) Select(f SelectFunc) *JSet {
    newSet := &JSet{ make([]*JEntry, 0, len(set.Entries)) }

    set.Each(func (entry *JEntry) {
        if f(entry) {
            newSet.Entries = append(newSet.Entries, entry)
        }
    })

    return newSet
}
