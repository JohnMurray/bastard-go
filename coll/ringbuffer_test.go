package coll

import (
	"testing"
)

//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/coll/ringbuffer.go.tmp -out=ringbuffer_string.go gen "Type=string"

func TestHasItemsEmpty(t *testing.T) {
	r := NewStringRingBuffer(2)
	if r.HasItems() {
		t.Fail()
	}
}

func TestHasItemsNonEmpty(t *testing.T) {
	r := NewStringRingBuffer(8)
	r.Push("one")
	if !r.HasItems() {
		t.Fatal("Does not have items after one push")
	}
	r.Pop()
	if r.HasItems() {
		t.Fatal("Buffer contains items after pop'ing all instances")
	}

	r.Push("two")
	if !r.HasItems() {
		t.Fatal("Does not have items after pop and second push")
	}
}

func TestPushOrder(t *testing.T) {
	r := NewStringRingBuffer(8)
	data := []string{"one", "two", "three"}
	for _, d := range data {
		r.Push(d)
	}

	for _, d := range data {
		val, err := r.Pop()
		if err != nil || d != val {
			t.Fatalf("Value should be '%s', was instead '%s'\n", d, val)
			return
		}
	}
}

func TestOverflow(t *testing.T) {
	r := NewStringRingBuffer(4)
	data := []string{"1", "2", "3", "4", "5"}
	for _, d := range data {
		r.Push(d)
	}

	for i := 0; i < 4; i++ {
		val, err := r.Pop()
		if err != nil || val == "5" {
			t.Fatalf("5 should not exist in the buffer")
			return
		}
	}
}

func TestUnderflow(t *testing.T) {

	r := NewStringRingBuffer(4)
	data := []string{"1", "2", "3", "4", "5"}
	for _, d := range data {
		r.Push(d)
	}

	for i := 0; i < 4; i++ {
		val, err := r.Pop()
		if err != nil || val == "5" {
			t.Fatalf("5 should not exist in the buffer")
			return
		}
	}
}
