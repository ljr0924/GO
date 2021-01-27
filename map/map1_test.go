package map

import "fmt"
import "testing"

func TestMap1(t *testing.T) {

    // 初始化map
    profile := map[string]string {
        "name": "github",
        "age": "18",
    }
    fmt.Printf("%+v\n", profile)
    for k, v := range profile {
        fmt.Printf("key: %s value: %s\n", k, v)
    }

    // 使用make初始化
    profile_make := make(map[string]string, 10)
    profile_make["name"] = "make"
    profile_make["age"] = "18"
    fmt.Printf("%+v\n", profile_make)

    // 增删改查
    profile_make["sex"] = "1"
    fmt.Printf("增加sex属性后，%+v\n", profile_make)

    delete(profile_make, "sex")
    fmt.Printf("删除sex属性后，%+v\n", profile_make)

    profile_make["name"] = "modify"
    fmt.Printf("修改name属性后，%+v\n", profile_make)

    fmt.Printf("profile_make name: %s\n", profile_make["name"])

    fmt.Println("查询map里面是否存在key")
    if v, ok := profile_make["notExist"]; !ok {
        fmt.Println("不存在key为notExists的值")
    } else {
        fmt.Printf("存在key为notExists的值, value: %s\n", v)
    }

    // map长度
    fmt.Printf("map 长度： %d, 容量：%d", len(profile_make))

}