package map_demo

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
    profileMake := make(map[string]string, 10)
    profileMake["name"] = "make"
    profileMake["age"] = "18"
    fmt.Printf("%+v\n", profileMake)

    // 增删改查
    profileMake["sex"] = "1"
    fmt.Printf("增加sex属性后，%+v\n", profileMake)

    delete(profileMake, "sex")
    fmt.Printf("删除sex属性后，%+v\n", profileMake)

    profileMake["name"] = "modify"
    fmt.Printf("修改name属性后，%+v\n", profileMake)

    fmt.Printf("profile_make name: %s\n", profileMake["name"])

    fmt.Println("查询map里面是否存在key")
    if v, ok := profileMake["notExist"]; !ok {
        fmt.Println("不存在key为notExists的值")
    } else {
        fmt.Printf("存在key为notExists的值, value: %s\n", v)
    }

    // map长度
    fmt.Printf("map 长度： %d, 容量：%d", len(profileMake), len(profileMake))

}