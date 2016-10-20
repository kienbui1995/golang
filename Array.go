package main

import "fmt"

func main() {
	fmt.Print
	//Khai bao mang
	// var mang [5]int
	// mang[4] = 100
	// fmt.Println(mang)
	fmt.Println("Nhap n: ")
	var n int
	fmt.Scanf("%f", &n)
	var mang [n]float64
	fmt.Println("Moi ban nhap so: ")
	for i := 0; i < n; i++ {
		fmt.Println("Nhap so thu", i)
		fmt.Scanf("%f", mang[i])
	}
	// Loi vi kich thuoc mang khong doi duoc phai dung Slice
}
