package model

import "fmt"

// Alphabet 是一个结构体，表示键值对数据
type Alphabet struct {
	Key   int
	Value string
}

// CreateAlphabetTable 创建一个字母表
func CreateAlphabetTable() (err error) {
	alphabetData := []Alphabet{
		{Key: 1, Value: "a"},
		{Key: 2, Value: "b"},
		{Key: 3, Value: "c"},
		{Key: 4, Value: "d"},
		{Key: 5, Value: "e"},
		{Key: 6, Value: "f"},
		{Key: 7, Value: "g"},
		{Key: 8, Value: "h"},
		{Key: 9, Value: "i"},
		{Key: 10, Value: "j"},
		{Key: 11, Value: "k"},
		{Key: 12, Value: "l"},
		{Key: 13, Value: "m"},
		{Key: 14, Value: "n"},
		{Key: 15, Value: "o"},
		{Key: 16, Value: "p"},
		{Key: 17, Value: "q"},
		{Key: 18, Value: "r"},
		{Key: 19, Value: "s"},
		{Key: 20, Value: "t"},
		{Key: 21, Value: "u"},
		{Key: 22, Value: "v"},
		{Key: 23, Value: "w"},
		{Key: 24, Value: "x"},
		{Key: 25, Value: "y"},
		{Key: 26, Value: "z"},
	}

	// 插入数据到数据库
	db.Create(&alphabetData)
	if db.Error != nil {
		fmt.Println("数据插入失败！", db.Error)
		return db.Error
	}
	fmt.Println("数据插入成功！")
	return nil
}
