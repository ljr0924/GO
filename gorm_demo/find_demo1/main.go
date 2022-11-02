package main

import (
	"errors"
	"fmt"
	. "go_demo/gorm_demo"
	"gorm.io/gorm"
)

func main() {

	// 获取第一条记录， 主键升序
	p1 := &Product{}
	DB.First(p1)
	fmt.Printf("p1:  %+v\n", p1)

	// 获取一条记录，不排序
	p2 := &Product{}
	DB.Where("id < 10 and id > 5").Take(&p2)
	fmt.Printf("p2:  %+v\n", p2)

	// 获取最后一条，主键降序 order by id desc
	p3 := &Product{}
	DB.Where("id < 10 and id > 5").Last(&p3)
	fmt.Printf("p3:  %+v\n", p3)

	p4 := &Product{}
	result := DB.Where("id=101").First(p4)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))

	p5 := map[string]interface{}{}
	result5 := DB.Table("products").Take(&p5)
	fmt.Println(result5.RowsAffected)
	fmt.Println(result5.Error)
	fmt.Printf("p5:  %+v\n", p5)

	p6 := &Product{}
	DB.Find(p6, "100").Take(&p5)
	fmt.Printf("p6:  %+v\n", p6)

	var pList []*Product
	DB.Where("id in ?", []int{1, 2, 3}).Select("id", "code", "price").Find(&pList)
	for _, v := range pList {
		fmt.Println(v)
	}

	p7 := &Product{}
	result7 := DB.Raw("select sum(price) as price from products where id = ?", "4").Scan(p7)
	fmt.Println(result7.Error)
	fmt.Printf("p7:  %+v\n", p7)

}
