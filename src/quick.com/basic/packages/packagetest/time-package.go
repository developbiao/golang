package packagetest

import (
	"fmt"
	"time"
)

func TestTime() {
	t := time.Now()
	fmt.Println(t)

	// 格式化时间
	s := t.Format("2006年1月2日 15:04:05")
	fmt.Println(s)

	// 字符串类型的时间
	str := "2020年9月1日"
	// 第一个参数是模板，第二个参数是要转换的时间字符串
	s1, _ := time.Parse("2006年1月2日", str)
	fmt.Println(s1)

	// 获取年月日信息
	year, month, day := time.Now().Date()
	fmt.Println(year, month, day)

	// 获取时分秒
	hour, minute, second := time.Now().Clock()
	fmt.Println(hour, minute, second)

	// 获取今年过了多少天了
	tday := time.Now().YearDay()
	fmt.Println(tday)

	// 获取今天是星期几
	weeday := time.Now().Weekday()
	fmt.Println(weeday, t.Local())
	fmt.Println(t.Zone())

}

func TestTimestamp() {
	t := time.Date(2008, 8, 8, 17, 30, 0, 0, time.UTC)
	timestamp := t.Unix()
	fmt.Println(timestamp)

	timestamp2 := time.Now().Unix()
	fmt.Println(timestamp2)

	timestamp3 := time.Now().UnixNano()
	fmt.Println(timestamp3)

	now := time.Now()
	t1 := now.Add(time.Minute * 3)

	// 计算两个时间间隔
	d := t1.Sub(now)
	fmt.Println(d)

	// 将指定的时间转为时间戳格式
	beforetime := "2020-08-15 00:00:00"
	timeLayout := "2006-01-02 15:04:05"
	loc := time.Now().Location()
	theTime, _ := time.ParseInLocation(timeLayout, beforetime, loc)
	aftertime := theTime.Unix()
	fmt.Println(theTime, aftertime)

	// 再将时间戳转换为日期
	dateTimeStr := time.Unix(aftertime, 0).Format(timeLayout)
	fmt.Println(dateTimeStr)



}