package main

import "fmt"

func main() {
	var m1 map[int]string
	var m2 = make(map[int]string)
	m3 := map[string]int{"英语": 99, "数学": 32, "语文": 91}

	fmt.Println(m1 == nil, m1)
	fmt.Println(m2 == nil, m2)
	fmt.Println(m3 == nil, m3)

	if m1 == nil {
		m1 = make(map[int]string)
	}

	m1[1] = "小猪"
	m1[2] = "小猫"

	val, ok := m1[2]
	fmt.Println(val, ok)

	m1[1] = "阿黄"
	fmt.Println(val, ok, len(m1))

	// Traversing map
	m4 := map[string]string{
		"name":    "ccmouse",
		"course":  "goalng",
		"site":    "google",
		"quality": "notbad",
	}

	for k, v := range m4 {
		fmt.Println(k, v)
	}

	map1 := make(map[string]string)
	map1["name"] = "张无忌"
	map1["sex"] = "男"
	map1["age"] = "21"
	map1["address"] = "明教"

	map2 := make(map[string]string)
	map2["name"] = "周芷若"
	map2["sex"] = "女"
	map2["age"] = "22"
	map2["address"] = "峨眉山"

	// save map to slice
	slice3 := make([]map[string]string, 0, 2)
	slice3 = append(slice3, map1)
	slice3 = append(slice3, map2)

	// traversing slice3
	for key, val := range slice3 {
		fmt.Println(key, val)
	}

}
