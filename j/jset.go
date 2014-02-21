package j

import (
    "io/ioutil"
    "os"
    "sort"
    "strconv"
    "strings"
    "time"
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
    for _, e := range set.Entries {
        f(e)
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

func (set *JSet) Len() int {
    return len(set.Entries)
}

func (set *JSet) Swap(i, j int) {
    set.Entries[i], set.Entries[j] = set.Entries[j], set.Entries[i]
}

func (set *JSet) Less(i, j int) bool {
    return set.Entries[i].Frecency > set.Entries[j].Frecency
}

func (set *JSet) Sort() *JSet {
    sort.Sort(set)

    return set
}

func (set *JSet) Limit(limit int) *JSet {
    set.Entries = set.Entries[0:limit]

    return set
}

func (set *JSet) Add(pathToAdd string) {
    // ignore home directory
    if pathToAdd == os.Getenv("HOME") {
        return
    }

    now := int(time.Now().Unix())

    // if entry already exists
    if entry, found := set.findByPath(pathToAdd); found {
        // add 1 to rank
        entry.Rank += 1

        // set timestamp to now
        entry.Timestamp = now
    } else {
        newEntry := NewJEntry(pathToAdd, 2, now)

        set.Entries = append(set.Entries, newEntry)
    }
}

func (set *JSet) findByPath(pathToFind string) (existing *JEntry, found bool) {
    existing = nil
    found    = false

    set.Each(func (entry *JEntry) {
        if entry.Path == pathToFind {
            existing = entry
            found    = true
        }
    })

    return existing, found
}

func (set *JSet) totalRankCount() (count float64) {
    set.Each(func (entry *JEntry) {
        count += entry.Rank
    })

    return count
}

func (set *JSet) age() {
    // get rank total
    totalRank := set.totalRankCount()

    // if rank total > 6000
    if totalRank > 6000 {
        // multiply all ranks by 0.99
        set.Each(func (entry *JEntry) {
            entry.Rank *= 0.99
        })
    }

    // reject all entries with rank < 1
    set = set.Select(func (entry *JEntry) bool {
        return entry.Rank >= 1
    })
}

func (set *JSet) writeToFile(file string) error {
    f, err := os.Create(file)
    defer f.Close()

    if err != nil {
        return err
    }

    for _, entry := range set.Entries {
        _, err := f.WriteString(entry.DataString() + "\n")

        if err != nil {
            return err
        }
    }

    f.Sync()

    return nil
}
