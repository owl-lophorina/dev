package main

import (
	"fmt"
	"strings"
)

func Question01() {
	//fmt.Println("callQ01")
	for i := 10; ; i++ {
		decimal := fmt.Sprintf("%d", i) // 10進数を文字列に変換
		binary := fmt.Sprintf("%b", i)  //2進数に変換し、文字列型に変換
		octal := fmt.Sprintf("%o", i)   //8進数に変換し、文字列型に変換

		if decimal == reverse(decimal) && binary == reverse(binary) && octal == reverse(octal) {
			fmt.Println(decimal)
			break
		}
	}
}

func reverse(str string) string {
	//逆順にした文字列を戻す
	slice := strings.Split(str, "")
	len := len(slice)
	for i := 0; i < len/2; i++ {
		slice[i], slice[len-i-1] = slice[len-i-1], slice[i]
	}
	return strings.Join(slice, "")
}

/*
どのようなロジックにするか？
・10～見つけるまでのfor文
・10進数を2進数と8進数に変換する
・それぞれに変換した数字を文字列に変換し、sliceに1文字ずつ格納する
・逆順にして、文字列に戻す
・逆順にする前後の文字列を比較して、3つ全て一致するならfor文を終了する
*/
