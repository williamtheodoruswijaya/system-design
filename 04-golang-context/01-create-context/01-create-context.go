package main

import (
	"context"
	"fmt"
)

func main() {
	// 1. Contoh membuat context Background
	background := context.Background()
	fmt.Println(background)

	/*
		context.Background() -> membuat empty context dimana semua attribute contextnya seperti:
			- Value
			- Done()
			- Timeout()
		itu masih kosong semua. Ini yang normalnya akan kita gunakan dan di passing ke function-function dari handler/controller layer (tempat context dibuat) ke service/usecase layer.
	*/

	// 2. Contoh membuat context TODO (biasa dipakai ketika kita masih ga jelas mau pakai context apa di function yang akan menerima context ini, mksd "context mana yang bakal di pakai" bakal dijelasin nanti)
	todo := context.TODO()
	fmt.Println(todo)
}
