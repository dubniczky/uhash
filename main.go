package main

import (
	"bufio"
	"errors"
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

func processInline(data string, hash hash.Hash) []byte {
    hash.Write([]byte(data))
    output := hash.Sum(nil)
    return output
}

func processStdin(hash hash.Hash) []byte {
    data := hashBuffer( bufio.NewReader(os.Stdin), hash )
    return data
}

func beginHash(args []string) (string, error) {
    // Load hashing method
    if len(os.Args) < 2 {
        return "", errors.New("Usage: uhash <method> <data>")
    }
    method := os.Args[1]
    hash := protocolSwitcher(method)
    if hash == nil {
        return "", errors.New(fmt.Sprintf("Unknown hashing method: %s\n", method))
    }

    // Hash from parameter
    if len(os.Args) > 2 {
        data := os.Args[2]
        return encodeHex(processInline(data, hash)), nil
    }

    // Hash from stdin
    return encodeHex(processStdin(hash)), nil
}

func main() {
    hash, err := beginHash(os.Args)
    if err != nil {
        stderr(err.Error())
        os.Exit(1)
    }
    fmt.Printf("%s\n", hash)
    os.Exit(0)
}
