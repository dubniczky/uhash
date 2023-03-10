package main

import (
	"crypto/sha1"
	"fmt"
)

func protocolSwitcher(method string) func([]byte) []byte {
    switch method {
        case "sha1":
            return hashSha1
        case "sha256":
            fmt.Println("sha256")
        case "sha512":
            fmt.Println("sha512")
    }
    return nil
}

func hashSha1(data []byte) []byte {
    h := sha1.New()
    h.Write(data)
	return h.Sum(nil)
}