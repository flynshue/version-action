package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMakeTagVersion(t *testing.T) {
	var (
		testUUID    = "0eab6508-5bac-11ed-9f17-c8f75077546c"
		testSha1Sum = "878eda878a3cb65a2d5bc9265b0820f5a61a3bff"
	)
	actualSha1Sum := sha1sum(testUUID)
	if actualSha1Sum != testSha1Sum {
		t.Fatalf("got %s, wanted: %s\n", actualSha1Sum, testSha1Sum)
	}
	if os.Getenv("GITHUB_OUTPUT") == "" {
		os.Setenv("GITHUB_OUTPUT", "/tmp/output")
	}
	if err := makeTagVersion(actualSha1Sum); err != nil {
		t.Fatal()
	}
	f, err := os.Open(os.Getenv("GITHUB_OUTPUT"))
	if err != nil {
		t.Fatal()
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
		os.Unsetenv("GITHUB_OUTPUT")
	}()
	r := bufio.NewReader(f)
	line, _, err := r.ReadLine()
	if err != nil {
		t.Fatal()
	}
	fmt.Println(string(line))
	if string(line) != getExpectedVersion() {
		t.Fatalf("got %s, wanted %s\n", line, getExpectedVersion())
	}
}

func getExpectedVersion() string {
	t := time.Now()
	return fmt.Sprintf("version=v%s-878e", strings.Fields(t.String())[0])
}
