package j

import (
    "fmt"
    "time"
)

type JEntry struct {
    Path      string
    Rank      float64
    Timestamp int
    Frecency  float64
}


func (entry *JEntry) String() string {
    return fmt.Sprintf("%s %f %d %f", entry.Path, entry.Rank, entry.Timestamp, entry.Frecency)
}

func NewJEntry(path string, rank float64, timestamp int) *JEntry {
    now := int(time.Now().Unix())

    dx := now - timestamp

    var frecency float64

    if dx < 3600 {
        frecency = rank * 4.0
    } else if dx < 86400 {
        frecency = rank * 2
    } else if dx < 604800 {
        frecency = rank / 2
    } else {
        frecency = rank / 4
    }

    return &JEntry{path, rank, timestamp, frecency}
}

