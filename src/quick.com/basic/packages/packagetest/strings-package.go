package packagetest

import (
	"fmt"
	"strings"
)

func TestStrings () {
	s1 := "ok let's go"
	// 是否包含指定字符串
	fmt.Println(strings.Contains(s1, "go"))

	// 是否包含指定的字符串中任意一个字符 有一个出现过就返回true
	fmt.Println(strings.ContainsAny(s1,"glass"))

	// 返回指定字符出现的次数
	fmt.Println(strings.Count(s1, "g"))

	// 文本的开头
	fmt.Println(strings.HasPrefix(s1, "ok"))

	// 文本的结尾
	fmt.Println(strings.HasSuffix(s1, ".php"))

	// 查找指定字符在字符串中存在的位置 如果不存在返回-1
	fmt.Println(strings.Index(s1, "b"))
	// 查找指定字符中任意一个字符出现在字符串中的位置
	fmt.Println(strings.IndexAny(s1, "s"))

	// 查找指定字符串出现在字符串中最后的一个位置
	fmt.Println(strings.LastIndex(s1, "s"))

	// 字符串拼接
	s2 := []string{"123n", "abc", "ss"}
	s3 := strings.Join(s2, "_")
	fmt.Println(s3)

	// 字符切割
	s4 := strings.Split(s3, "_")
	fmt.Println(s4)

	// 字符串替换
	s5 := "go is very good"
	s6 := strings.Replace(s5, "go", "php", -1)
	fmt.Println(s6)


}
