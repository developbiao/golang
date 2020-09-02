package main

import (
	"encoding/json"
	"fmt"
)

type Prescription struct {
	Name     string        `json:"name"` // 重新指定json字段为小写输出
	Unit     string        `json:"unit"`
	Additive *Prescription `json:"additive,omitempty"`
}

func structToJson() string {
	p := Prescription{}
	p.Name = "鹤顶红"
	p.Unit = "1.3kg"
	p.Additive = &Prescription{
		Name: "砒霜",
		Unit: "0.5kg",
	}

	buf, err := json.Marshal(p)
	if err != nil {
		fmt.Println("err = ", err)
		return ""
	}

	var json_str string = string(buf)
	return json_str

}

// json to struct
func jsonToStruct(jsonstr string) Prescription {
	var p Prescription
	if err := json.Unmarshal([]byte(jsonstr), &p); err != nil {
		fmt.Println(err)
	}
	return p
}

func main() {
	// struct to json
	json_str := structToJson()
	fmt.Println("json = ", json_str)

	// json to struct
	p := jsonToStruct(json_str)
	fmt.Println(p)

}
