package main

import "fmt"

func introduce(name string, age int) {
	fmt.Printf("你好，我叫 %s，今年 %d 岁，是一名前端开发者。\n", name, age)
}

func main() {
	fmt.Println("hello world")
	// 调用新函数示例
	introduce("Alice", 28)
}
