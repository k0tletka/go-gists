package main

import (
    "fmt"
    "strings"
    "bytes"

    "cryptoalgs-des/des"
)

func main() {
    // Some encrypting data
    reader := strings.NewReader("Hello, world!")
    key := []byte("P@ssw0r")

    buffer := &bytes.Buffer{}
    des.EncryptData(key, reader, buffer)

    fmt.Println(buffer.Bytes())
}
