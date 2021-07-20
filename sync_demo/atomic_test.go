package sync_demo

import (
    "sync/atomic"
    "testing"
)

func TestAddInt32(t *testing.T) {
    var i int32 = 1
    t.Log(i)
    atomic.AddInt32(&i, 2)
    t.Log(i)
}
