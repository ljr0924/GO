package struct_demo

import "testing"

type user struct {
    name string
    age  byte
    addr address
}

type address struct {
    country  string
    province string
    city     string
}

func TestInitStruct(t *testing.T) {
    u := user{
        name: "u1",
        age:  1,
        addr: address{
            country:  "china",
            province: "GuangDong",
            city:     "ShenZhen",
        },
    }
    t.Logf("%+v", u)

    u.addr.city = "GuangZhou"

    t.Logf("%+v", u)
}
