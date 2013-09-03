package j

import (
    "fmt"
    "regexp"
    s "strings"
)


func Match(set *JSet, query string) {
    regex := s.Join(s.Split(query, " "), "|")
    matcher := regexp.MustCompile(regex)

    matches := set.Select(func (entry *JEntry) bool {
        return matcher.MatchString(entry.Path)
    })

    matches.Each(func (entry *JEntry) {
        fmt.Println(entry)
    })
}
