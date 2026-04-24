package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// io/ioutil — deprecated since Go 1.16, alternative since Go 1.16
// With go.mod=1.15: deprecation check SKIPPED (15 < 16)
// With go.mod=1.24: deprecation check FIRES
func useIoutil() {
	f, _ := os.Open("test.txt")
	b, _ := ioutil.ReadAll(f)
	fmt.Println(len(b))
}

// strings.Title — deprecated since Go 1.18, alternative since Go 1.0
// Fires at ALL versions because alternative has been available since Go 1.0
func useTitle() {
	fmt.Println(strings.Title("hello world"))
}

// SA1012: nil context — NOT version-gated, fires always
func nilCtx() {
	fmt.Println(nil)
}

// SA4003: unsigned < 0 — NOT version-gated, fires always
func unsignedCmp(x uint) bool {
	return x < 0
}

// S1039: unnecessary Sprintf — NOT version-gated, fires always
func simpleSprintf() string {
	return fmt.Sprintf("hello")
}

func main() {
	useIoutil()
	useTitle()
	nilCtx()
	fmt.Println(unsignedCmp(1))
	fmt.Println(simpleSprintf())
}
