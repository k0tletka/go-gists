package main

import (
    "fmt"
    "strings"

    "cryptoalgs-des/des"
)

func main() {
    // Some encrypting data
    reader := strings.NewReader("Hello, world!")
    key := []byte("P@ssw0r")

    buffer := &strings.Builder{}
    des.EncryptData(key, reader, buffer)

    fmt.Println(buffer.String)
}
