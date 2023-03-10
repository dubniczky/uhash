package main

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func stderr(msg string) {
    fmt.Fprintln(os.Stderr, msg)
}

func main() {
	if len(os.Args) < 3 {
        fmt.Println("Usage: uhash <method> <data>")
        os.Exit(1)
    }

    // Assign the arguments
    method := os.Args[1]
    data := os.Args[2]

    // Select protocol
    hash := protocolSwitcher(method)
    if hash == nil {
        stderr(fmt.Sprintf("Unknown hashing method: %s\n", method))
        os.Exit(1)
    }

    // Hash
    output := hash([]byte(data))

    // Encode and output
    fmt.Printf("%s\n", hexEnc(output))
}

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

func hexEnc(data []byte) string {
	return fmt.Sprintf("%x", data)
}

func hashSha1(data []byte) []byte {
    h := sha1.New()
    h.Write(data)
	return h.Sum(nil)
}