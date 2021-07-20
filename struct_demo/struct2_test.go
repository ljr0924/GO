package struct_demo

import (
    "encoding/json"
    "testing"
)

type Profile struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func TestStructTag(t *testing.T) {
    jsonText := "{\"name\":\"u1\",\"age\":1}"

    var p Profile

    _ = json.Unmarshal([]byte(jsonText), &p)

    t.Logf("%+v", p)


}
