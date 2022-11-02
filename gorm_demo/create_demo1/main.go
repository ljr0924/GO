package main

import (
	. "go_demo/gorm_demo"
	"strconv"
)

func main() {

	// 迁移 schema
	err := DB.AutoMigrate(&Product{})
	if err != nil {
		panic("failed migrate tables")
		return
	}

	var pList []*Product
	for i := 1; i <= 100; i++ {
		pList = append(pList, &Product{Code: strconv.Itoa(i), Price: uint(i)})
	}
	DB.Create(&pList)

}
