package main

import (
    "log"
    "bufio"
    "net"
    "os"
    "fmt"

    ln "l2-clientserver/linknetwork"
)

func usage() {
    fmt.Printf("Usage: %s interface server_mac_address\n", os.Args[0])
    os.Exit(1)
}

func main() {
    if len(os.Args) != 3 {
        usage()
    }

    interfaceName, serverMACAddress := os.Args[1], os.Args[2]

    mac, err := net.ParseMAC(serverMACAddress)

    if err != nil {
        panic(err)
    }

    c, err := ln.DialLink(interfaceName, ln.LinkAddr{mac})

    if err != nil {
        log.Fatalln(err)
    }

    reader := bufio.NewReader(os.Stdin)

    for {
        input, err := reader.ReadString('\n')

        if err != nil {
            log.Fatalln(err)
        }

        if _, err := c.Write([]byte(input)); err != nil {
            log.Fatalln(err)
        }
    }
}
