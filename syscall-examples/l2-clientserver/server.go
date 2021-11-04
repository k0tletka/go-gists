package main

import (
    "log"
    "fmt"
    "os"

    ln "l2-clientserver/linknetwork"
)

func usage() {
    fmt.Printf("Usage: %s interface\n", os.Args[0])
    os.Exit(1)
}

func main() {
    if len(os.Args) != 2 {
        usage()
    }

    interfaceName := os.Args[1]
    c, err := ln.Listen(interfaceName)

    if err != nil {
        log.Fatalln(err)
    }
    defer c.Close()

    readBuffer := make([]byte, ln.LinkMTU)

    for {
        var n int
        var err error

        if n, err = c.Read(readBuffer); err != nil {
            log.Fatalln(err)
        }

        log.Println("Got message: ", string(readBuffer[:n]))
    }
}
