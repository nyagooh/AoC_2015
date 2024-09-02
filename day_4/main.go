package main

import (
    "crypto/md5"
    "fmt"
    "encoding/hex"
)

func main() {
    secretKey := "bgvyzdsv"
    number := 1

    for {
        input := fmt.Sprintf("%s%d", secretKey, number)
        hash := md5.Sum([]byte(input))
        hashHex := hex.EncodeToString(hash[:])
        if hashHex[:5] == "00000" {
            fmt.Println( number)
            break
        }
        number++
    }
}