package redis_demo

import (
    "testing"

    "go_demo/redis_demo/client"
)

var c = client.Client

func TestBitmap(t *testing.T) {

    var err error

    // mybit 00000100
    err = c.SetBit("mybit1", 5, 1).Err()
    if err != nil {
        t.Fatal(err)
    }

    r, err := c.GetBit("mybit1", 5).Result()
    if err != nil {
        t.Fatal(err)
    }
    t.Log(r)

    // mybit 01000100
    err = c.SetBit("mybit1", 2, 1).Err()
    if err != nil {
        t.Fatal(err)
    }

    r, err = c.BitCount("mybit1", nil).Result()
    if err != nil {
        t.Fatal(err)
    }
    t.Log(r)

    // mybit 10110000
    c.SetBit("mybit2", 0, 1)
    c.SetBit("mybit2", 2, 1)
    c.SetBit("mybit2", 3, 1)

    // 逻辑并
    // 101100
    // 001001
    c.BitOpAnd("destkeyAND", "mybit1", "mybit2")
    r, err = c.BitCount("destkeyAND", nil).Result()
    if err != nil {
        t.Fatal(err)
    }
    t.Log("AND ", r)

    // 逻辑或
    // 101100
    // 001001
    c.BitOpOr("destkeyOR", "mybit1", "mybit2")
    r, err = c.BitCount("destkeyOR", nil).Result()
    if err != nil {
        t.Fatal(err)
    }
    t.Log("OR ", r)

    // 逻辑异或
    // 101100
    // 001001
    c.BitOpXor("destkeyXOR", "mybit1", "mybit2")
    r, err = c.BitCount("destkeyXOR", nil).Result()
    if err != nil {
        t.Fatal(err)
    }
    t.Log("XOR ", r)

    // 逻辑异或
    // 00100100
    c.BitOpNot("destkeyNOT", "mybit1")
    r, err = c.BitCount("destkeyNOT", nil).Result()
    if err != nil {
        t.Fatal(err)
    }
    t.Log("NOT ", r)

    // 逻辑异或
    // 10110000
    c.BitOpNot("destkeyNOT", "mybit2")
    r, err = c.BitCount("destkeyNOT", nil).Result()
    if err != nil {
        t.Fatal(err)
    }
    t.Log("NOT ", r)

    c.FlushDb()

}
