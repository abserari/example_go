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

type ping struct {
	X, Y, Z int
	Name    string
	Args    interface{}
}

type pong struct {
	X, Y *int32
	Name string
	Args interface{}
}

type args struct {
	A, B int
}

func main() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	gob.RegisterName("two", args{})
	err := enc.Encode(ping{3, 4, 5, "Pythagoras", args{2, 3}})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	var pong pong
	err = dec.Decode(&pong)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	fmt.Println(pong)
	//if value is a pointer, must aasert first or if is nil
	//or get runtime error
	fmt.Println(*pong.X)
}
