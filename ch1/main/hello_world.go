package main

import (
	"fmt"
	"os"
)

/*
 * 1. main函数不支持任何返回值
 * 2. os.Exit()返回值
 * 3. os.Args 获取命令行参数
 */
func main() {
	fmt.Println(os.Args, "123") // [C:\Users\wallenhuang\AppData\Local\Temp\GoLand\___go_build_hello_world_go.exe]
	fmt.Println("长度:", len(os.Args))
	//go run hello_world.go 我叫黄飞龙 我是李圭一 我叫张三, 打印结果:Hello World, 我叫黄飞龙
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
	fmt.Println("Hello World!")
	os.Exit(-1)
}
