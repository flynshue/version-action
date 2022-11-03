package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"crypto/sha1"

	"github.com/google/uuid"
)

func main() {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatalln(err)
	}
	hash := sha1sum(u.String())
	if err := makeTagVersion(hash); err != nil {
		log.Fatalln(err)
	}
}

func sha1sum(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	sha1Sum := h.Sum(nil)
	return fmt.Sprintf("%x", sha1Sum)
}

func makeTagVersion(hash string) error {
	t := time.Now()
	date := strings.Fields(t.String())[0]
	version := fmt.Sprintf("v%s-%s", date, hash[0:4])
	outFile := os.Getenv("GITHUB_OUTPUT")
	f, err := os.OpenFile(outFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("version=%s\n", version))
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
