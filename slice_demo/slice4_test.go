package slice_demo

import "testing"

func TestAppendSlice(t *testing.T) {

    s := make([]int, 0, 1)

    // 超出容量之后，会重新分配数组内存
    // 容量会成倍增长，对于较大的切片，会尝试扩容1/4并非2倍
    s1 := append(s, 1)
    s2 := append(s, 1, 2)
    s3 := append(s, 1, 2, 3)
    t.Logf("s内存地址：%p，长度：%d，容量：%d", s, len(s), cap(s))
    t.Logf("s1内存地址：%p，长度：%d，容量：%d", s1, len(s1), cap(s1))
    t.Logf("s2内存地址：%p，长度：%d，容量：%d", s2, len(s2), cap(s2))
    t.Logf("s3内存地址：%p，长度：%d，容量：%d", s3, len(s3), cap(s3))

}

func TestCopy(t *testing.T) {

    s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

    s1 := s[5:8]
    n := copy(s[4:], s1) // 在同一底层数组的不同区间复制
    t.Log(n, s)

    // 两个切片对象间复制数据，允许指向同一底层数组，允许目标区间重叠。
    // 最终所复制长度以较短的切片长度（len）为准
    s2 := make([]int, 6)
    n = copy(s2, s)
    t.Log(n, s2)

}

type Data struct {
    List []int
}

func TestAppend(t *testing.T) {

    d := Data{}

    for i := 0; i < 10; i++ {
        d.List = append(d.List, i)
    }

    t.Log(d)
}
