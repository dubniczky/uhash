package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"strings"
)

func protocolSwitcher(method string) func([]byte) []byte {
	method = strings.ToLower(method)
    switch method {
        case "sha1", "sha-1":
            return hashSha1
        case "sha256", "sha-256":
            return hashSha256
        case "sha512", "sha-512":
            return hashSha512
        case "md5":
            return hashMd5
    }
    return nil
}

func hashSha1(data []byte) []byte {
    h := sha1.New()
    h.Write(data)
	return h.Sum(nil)
}

func hashSha256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func hashSha512(data []byte) []byte {
	h := sha512.New()
	h.Write(data)
	return h.Sum(nil)
}

func hashMd5(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}
