package j

import (
  "fmt"
  "time"
)

type JEntry struct {
  Path         string
  Rank         float64
  Timestamp    int
  WeightedRank float64
}


func (entry *JEntry) String() string {
    return fmt.Sprintf("Entry: %s %f %d %f", entry.Path, entry.Rank, entry.Timestamp, entry.WeightedRank)
}

func NewJEntry(path string, rank float64, timestamp int) *JEntry {
    now := int(time.Now().Unix())

    dx := now - timestamp

    var weightedRank float64

    if dx < 3600 {
        weightedRank = rank * 4.0
    } else if dx < 86400 {
        weightedRank = rank * 2
    } else if dx < 604800 {
        weightedRank = rank / 2
    } else {
        weightedRank = rank / 4
    }

    return &JEntry{path, rank, timestamp, weightedRank}
}

