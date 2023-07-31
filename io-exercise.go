package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	shells, err := os.ReadFile("/etc/shells")
	if err != nil { // エラーかどうか
		fmt.Println(err)
		return
	}
	err = os.WriteFile("shells.txt", shells, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(shells)) // stringに変換
	fmt.Println("完了")
}