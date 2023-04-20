package main

import (
	"errors"
	"fmt"
	"os"
)

func StrUnpacking(s string) (string, error) {
	const zeroUnicode, nineUnicode = 48, 57
	var result string = ""
	//нужно, чтобы отключать проверку на число
	//и чтобы отлючать проверку на `\`
	var skipNext bool
	for i, r := range s {
		//проверка на `\`
		if r == 92 && !skipNext {
			skipNext = true
			continue
		}

		//сколько добавить повторений в строку result
		countRep := 1
		//проверка на число
		if r >= zeroUnicode && r <= nineUnicode && !skipNext {
			//проверка ошибки на первый элемент строки
			if i == 0 {
				return "", errors.New("string cannot start with a number")
			}
			//вычетаем 48, чтобы привести число к 10-ой системе
			//и -1, потому что в предыдущей итерации уже добавился
			//этот же символ
			countRep = int(r) - zeroUnicode - 1
			//получаем предыдущий символ
			r = rune(s[i-1])
		}
		for i := 0; i < countRep; i++ {
			result += string(r)
		}
		skipNext = false
	}
	return result, nil
}

func main() {
	strs := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		`qwe9\9`,
		`qwe\45`,
		`qwe\\5`,
	}

	for i := range strs {
		unpacked, err := StrUnpacking(strs[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(unpacked)
	}
}
