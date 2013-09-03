package j

import "fmt"

func Dump(set *JSet) error {
    for _, entry := range set.Entries {
        fmt.Println(entry)
    }

    return nil
}
