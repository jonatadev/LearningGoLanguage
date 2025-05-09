package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
)

type Message struct {
	Content string
	once    sync.Once
	sig     string //cached signature
}

func main() {

	m := Message{
		Content: "There is nothing more deceptive than an obvious fact."}
	fmt.Println(m.Signature())
	fmt.Println(m.Signature())
}

// Signature returns the digital signature of the message
func (m *Message) Signature() string {
	m.once.Do(m.calcSig)
	return m.sig
}

func (m *Message) calcSig() {
	log.Printf("Calculating signature")
	h := sha1.New()
	io.Copy(h, strings.NewReader(m.Content))
	m.sig = fmt.Sprintf("%x", h.Sum(nil))
}
