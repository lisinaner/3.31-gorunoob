package main

import "fmt"

type Books1 struct {
   title string
   author string
   subject string
   book_id int
}


func main1() {

    // 创建一个新的结构体
    fmt.Println(Books1{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    fmt.Println(Books1{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

    // 忽略的字段为 0 或 空
   fmt.Println(Books1{title: "Go 语言", author: "www.runoob.com"})
}