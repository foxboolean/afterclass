package main

import (
	"fmt"
	"testing"
)

func TestMessage_Decode(t *testing.T) {
	m := &message{
		ver: 1,
		op: 2,
		seq: 3,
		body: []byte("test websocket"),
	}
	msg := m.Encode()

	dm := &message{}
	err := dm.Decode(msg)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(dm.body))
}