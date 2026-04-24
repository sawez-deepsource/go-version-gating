package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ============================================================
// VERSION-GATED: SA1003 — bool to binary.Write
// Fires ONLY for go.mod < 1.8 (bool wasn't supported before 1.8)
// ============================================================
func binaryBool() {
	binary.Write(os.Stdout, binary.LittleEndian, true)
}

// ============================================================
// VERSION-GATED: S1016 — struct type conversion
// Fires ONLY for go.mod >= 1.8 (tags ignored since 1.8)
// ============================================================
type T1 struct {
	A int `+"`"+`json:"a"`+"`"+`
}
type T2 struct {
	A int `+"`"+`json:"b"`+"`"+`
}

func structConvert() {
	x := T1{A: 1}
	y := T2{A: x.A}
	fmt.Println(y)
}

// ============================================================
// VERSION-GATED: GO-W1009 — deprecated ioutil (since 1.16)
// ============================================================
func deprecatedIoutil() {
	f, _ := os.Open("test.txt")
	b, _ := ioutil.ReadAll(f)
	fmt.Println(len(b))
}

// ============================================================
// VERSION-GATED: GO-W1009 — deprecated strings.Title (since 1.18)
// ============================================================
func deprecatedTitle() {
	fmt.Println(strings.Title("hello world"))
}

// ============================================================
// NOT VERSION-GATED: SA4003 — unsigned < 0 (fires always)
// ============================================================
func unsignedCmp(x uint) bool {
	return x < 0
}

// ============================================================
// NOT VERSION-GATED: S1039 — unnecessary Sprintf (fires always)
// ============================================================
func simpleSprintf() string {
	return fmt.Sprintf("hello")
}

func main() {
	binaryBool()
	structConvert()
	deprecatedIoutil()
	deprecatedTitle()
	fmt.Println(unsignedCmp(1))
	fmt.Println(simpleSprintf())
}
