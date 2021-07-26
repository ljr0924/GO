package json_demo

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {

	a := `{"gender":1}`
	b := make(map[string]interface{})
	_ = json.Unmarshal([]byte(a), &b)



}