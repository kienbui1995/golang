package main

import "fmt"

func main() {
	//Mot so ham xu ly voi String
	fmt.Println(len("Hello World"))
	fmt.Println("HNllo, World"[1])
	// Output là 101, do nó trả về kiểu Int . 101 tương ứng với chữ e trong bảng mã Ascii
	fmt.Println("Hello, " + "World")
}
