package redis_demo

import (
    "math/rand"
    "strconv"
    "testing"
)

func TestTag(t *testing.T) {

    var tag string

    // 构造数据
    // 100个标签
    cusCnt := 100000
    for i := 1; i <= 100; i++ {
        tag = "tag_" + strconv.Itoa(i)
        t.Logf("%s start", tag)
        cnt := 0
        for cnt < cusCnt {
            index := int64(rand.Intn(cusCnt))
            c.SetBit(tag, index, 1)
            cnt++
        }
        r, _ := c.BitCount(tag, nil).Result()
        t.Logf("%s customer %d", tag, r)
    }



}

// 获取两个标签并集
func TestTagOr(t *testing.T) {
    c.BitOpOr("ret1", "tag_1", "tag_2")
    r, _ := c.BitCount("ret1", nil).Result()
    t.Log("ret1 ", r)

    c.Del("ret1")
}

// 获取两个标签交集
func TestTagAnd(t *testing.T) {
    c.BitOpAnd("ret2", "tag_1", "tag_2")
    r, _ := c.BitCount("ret2", nil).Result()
    t.Log("ret2 ", r)

    c.Del("ret2")

}

// 仅包含标签1或标签2
// 连续异或不符合 仅包含条件
func TestTagXOR(t *testing.T) {
    c.BitOpXor("ret3", "tag_2", "tag_1")
    r, _ := c.BitCount("ret3", nil).Result()
    t.Log("ret3 ", r)
}

// 属于标签1但不属于标签2的客户
func TestTag1(t *testing.T) {

    c.BitOpNot("tag_2_NOT", "tag_2")
    c.BitOpAnd("ret4", "tag_1", "tag_2_NOT")
    r, _ := c.BitCount("ret4", nil).Result()
    t.Log("ret4 ", r)

    c.Del("ret3", "ret4")
}
