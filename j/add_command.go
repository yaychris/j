package j

func Add(set *JSet, pathToAdd string, file string) error {
    set.Add(pathToAdd)
    set.age()

    return set.writeToFile(file)
}
