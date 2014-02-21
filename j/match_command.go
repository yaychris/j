package j

import (
    "fmt"
    "regexp"
    s "strings"
)


func Match(set *JSet, query string, matchLimit int) {
    regex   := s.Join(s.Split(query, " "), "|")
    matcher := regexp.MustCompile(regex)

    matches := set.Select(func (entry *JEntry) bool {
        return matcher.MatchString(entry.Path)
    }).Sort()

    if matchLimit > 0 {
        if matchLimit > matches.Len() {
            matchLimit = matches.Len()
        }

        matches.Limit(matchLimit)
    }

    matches.Each(func (entry *JEntry) {
        if matches.Len() > 1 {
            fmt.Println(entry.Path)
        } else {
            fmt.Print(entry.Path)
        }
    })
}
