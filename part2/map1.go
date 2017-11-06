package main

import "fmt"

func main() {
	/* 创建 map */
	countryCapitalMap := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}

	fmt.Println("原始 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 删除元素 */
	delete(countryCapitalMap, "France1")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	for k, v := range countryCapitalMap { //key和value
		fmt.Println(k, v)
	}
	c, b := countryCapitalMap["France"] //第一个参数是value 第二个参数是看有没有这个key 返回一个Boolean
	fmt.Println(c, b)
}
