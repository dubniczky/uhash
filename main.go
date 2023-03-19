package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/jzelinskie/whirlpool"
)

func stderr(msg string) {
    fmt.Fprintln(os.Stderr, msg)
}

func encodeHex(data []byte) string {
	return fmt.Sprintf("%x", data)
}

func protocolSwitcher(method string) hash.Hash {
	method = strings.ToLower(method)
    switch method {
        case "sha1", "sha-1":
            return sha1.New()
		case "sha224", "sha-224":
			return sha512.New512_224()
        case "sha256", "sha-256":
            return sha256.New()
		case "sha384", "sha-384":
			return sha512.New384()
        case "sha512", "sha-512":
            return sha512.New()
		case "whirlpool":
			return whirlpool.New()
        case "md5":
            return md5.New()
    }
    return nil
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
