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
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
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
        // SHA 1
        case "sha-1", "sha1":
            return sha1.New()

        // SHA 2
		case "sha-224", "sha224":
			return sha512.New512_224()
        case "sha-256", "sha256":
            return sha256.New()
		case "sha-384", "sha384":
			return sha512.New384()
        case "sha-512", "sha512":
            return sha512.New()

        // SHA 3
        case "sha3-224", "sha3224":
			return sha3.New224()
        case "sha3-256", "sha3256":
			return sha3.New256()
        case "sha3-384", "sha3384":
			return sha3.New384()
        case "sha3-512", "sha3512":
            return sha3.New512()

        // MD
        case "md4":
            return md4.New()
        case "md5":
            return md5.New()
        case "ripemd-160", "ripemd160":
            return ripemd160.New()

        // Others
        case "whirlpool":
			return whirlpool.New()
    }
    return nil
}

func hashBuffer(r *bufio.Reader, hash hash.Hash) []byte {
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
