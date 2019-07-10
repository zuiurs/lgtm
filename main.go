package main

import (
	"math/rand"
	"os"
	"time"
)

var lgtm = []byte{'L', 'G', 'T', 'M'}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for i := 0; i < len(lgtm); i++ {
		os.Stderr.Write([]byte{lgtm[rand.Intn(len(lgtm))]})
	}
	os.Stderr.Write([]byte{'\n'})
}
