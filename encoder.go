package main

import "fmt"

func encodeHex(data []byte) string {
	return fmt.Sprintf("%x", data)
}
