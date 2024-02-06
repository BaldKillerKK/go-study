package main

import "fmt"

func main() {
	// map初始化-字面量初始化
	m1 := map[string]string{
		"a":         "b",
		"changeMap": "值传递",
	}
	// map初始化-make初始化
	m2 := make(map[string]string, 8)

	// 没有给map赋值 默认值为nil
	var m3 map[string]string
	// panic: assignment to entry in nil map
	m3["a"] = "b"

	// map赋值
	m1["c"] = "d"
	m2["a"] = "b"

	// map修改
	m1["a"] = "change"
	m2["a"] = "change"

	// map删除
	delete(m1, "c")
	delete(m2, "a")

	changeMap(m1)

	// map遍历
	for k, v := range m1 {
		fmt.Printf("key : %s,value : %s\r\n", k, v)
	}
}

// map传递是引用传递
func changeMap(m map[string]string) {
	// 判断获取的key是否存在,如果key存在就修改
	if _, exist := m["changeMap"]; exist {
		m["changeMap"] = "引用传递"
	}
}
