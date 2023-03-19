package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"strings"

	"github.com/jzelinskie/whirlpool"
)

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
