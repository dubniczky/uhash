package main

import (
	"bufio"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func stderr(msg string) {
    fmt.Fprintln(os.Stderr, msg)
}

func hashBuffer(r *bufio.Reader , hash hash.Hash) []byte {
    nBytes, nChunks := int64(0), int64(0)
    buf := make([]byte, 0, 4*1024)
    for {
        n, err := r.Read(buf[:cap(buf)])
        buf = buf[:n]
        if n == 0 {
            if err == nil {
                continue
            }
            if err == io.EOF {
                break
            }
            log.Fatal(err)
        }
        nChunks++
        nBytes += int64(len(buf))
        
        hash.Write(buf)
        
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
    }
    stderr(fmt.Sprintf("Read %d bytes from stdin", nBytes))
    return hash.Sum(nil)
}

func main() {
    // Load hashing method
    if len(os.Args) < 2 {
        fmt.Println("Usage: uhash <method> <data>")
        os.Exit(1)
    }
    method := os.Args[1]
    hash := protocolSwitcher(method)
    if hash == nil {
        stderr(fmt.Sprintf("Unknown hashing method: %s\n", method))
        os.Exit(1)
    }

    // Hash from parameter
    if len(os.Args) > 2 {
        data := os.Args[2]
        hash.Write([]byte(data))
        output := hash.Sum(nil)
        fmt.Printf("%s\n", encodeHex(output))
        os.Exit(0)
    }

    // Hash from stdin
    data := hashBuffer( bufio.NewReader(os.Stdin), protocolSwitcher(method) )
    fmt.Printf("%s\n", encodeHex(data))
    os.Exit(0)
}
