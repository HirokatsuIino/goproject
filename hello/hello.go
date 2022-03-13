package main

import "fmt"

// ２つの引数を足す関数
func add(x int, y int) int { return x + y }
// 2つの引数を掛ける関数
func mul(x, y int) int { return x * y}
// 文字を入れ替える関数
func swap(x, y string)(string, string){ return y, x }

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println(add(3, 4)) // 7
	fmt.Println(mul(3, 4)) // 12
	fmt.Println(swap("first", "second")) // second first
}
