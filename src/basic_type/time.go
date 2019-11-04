package main

import (
    "fmt"
    "time"
)

var week time.Duration

func main() {
    t := time.Now()
    fmt.Println(t)
    fmt.Println("%4d-%02d-%02d\n", t.Year(), t.Month(), t.Day())

    t = time.Now().UTC()
    fmt.Println(t)

    // calculating times:
    week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
    week_from_now := t.Add(time.Duration(week))
    fmt.Println(week_from_now)

    // formatting times
    fmt.Println(t.Format(time.RFC822))
    fmt.Println(t.Format(time.ANSIC))
    fmt.Println(t.Format("02 Jan 2008 22:59"))
    s := t.Format("20030604")
    fmt.Println(t, "=>", s)
}




