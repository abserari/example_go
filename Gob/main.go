/*
this is  for gob example

to exercise encode and Decode

*/
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Ping struct {
	X, Y, Z int
	Name    string
	Args    interface{}
}

type Pong struct {
	X, Y *int32
	Name string
	Args interface{}
}

type Args struct {
	A, B int
}

func main() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	gob.RegisterName("two", Args{})
	err := enc.Encode(Ping{3, 4, 5, "Pythagoras", Args{2, 3}})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	var pong Pong
	err = dec.Decode(&pong)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	fmt.Println(pong)
	//if value is a pointer, must aasert first or if is nil
	//or get runtime error
	fmt.Println(*pong.X)
}
