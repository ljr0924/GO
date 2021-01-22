package main

import "sync"
import "testing"

var mtx sync.Mutex

func call() {
	mtx.Lock()
	mtx.Unlock()
}

func deferCall() {
	mtx.Lock()
	defer mtx.Unlock()
}

func BenchmarkCall(b *testing.B) {
	call()
}

func BenchmarkDeferCall(b *testing.B) {
	deferCall()
}