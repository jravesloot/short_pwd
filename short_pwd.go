 package main

 import (
     "fmt"
     "os"
     "strings"
)

const (
    maxPathLength int = 30
    maxWordLength int = 2
)

func main() {
    fmt.Println(abbreviatePath())
}

// abbreviates User's home dir to '~' if part of path, and any directories in
// path with - or _ separated segments longer than maxWordLength
func abbreviatePath() string {
    // TODO check err?
    path, _ := os.Getwd()
    userHome, homeIsSet := os.LookupEnv("HOME") // portable?
    if homeIsSet {
        path = abbreviateHome(path, userHome)
    }
    if len(path) <= maxPathLength {
        return path
    }
    path = _abbreviatePath(path)
    return path
}

// if home directory present in path - eg. /Users/username - replace
// with '~' abbreviation
func abbreviateHome(path, userHome string) string {
    if strings.HasPrefix(path, userHome) {
        return strings.Replace(path, userHome, `~`, 1)
    }
    return path
}

// path abbreviation algorithm, operating on entire path as string
func _abbreviatePath(path string) string {
    var letterCount int
    newPathSlice := make([]byte, len(path))
    pathBytes := make([]byte, len(path))
    pathReader := strings.NewReader(path)
    numRead, err := pathReader.Read(pathBytes)

    if err != nil {
        return fmt.Sprintf("%v - %s", err, path)
    }

    for i := 0; i < numRead; i++ {
        switch pathBytes[i] {
        case '/', '_', '-':
            letterCount = 0
            newPathSlice = append(newPathSlice, pathBytes[i])
        default:
            if letterCount > maxWordLength {
                continue
            } else if letterCount == maxWordLength {
                letterCount++
                newPathSlice = append(newPathSlice, '.')
            } else {
                letterCount++
                newPathSlice = append(newPathSlice, pathBytes[i])
            }
        }
    }
    return string(newPathSlice)
}

