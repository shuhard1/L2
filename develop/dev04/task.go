package main

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {

	lenS := len(s)
	lenT := len(t)
	if lenS != lenT {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {
		anagramMap[string(s[i])]++
	}

	for i := 0; i < lenT; i++ {
		anagramMap[string(t[i])]--
	}

	for i := 0; i < lenS; i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}

	return true
}

// возавращает массив без повторов
func getUnique(slice []string) []string {
	resMap := make(map[string]struct{})
	result := []string{}

	//ключ не может повторятся, поэтому
	//так можно получить только уникальные элементы в слайсе
	for _, key := range slice {
		resMap[key] = struct{}{}
	}

	//записываем ключи resMap в слайс
	for key := range resMap {
		result = append(result, key)
	}
	return result
}

func searchAnagram(dictionary []string) map[string][]string {
	anagrams := make(map[string][]string)
	//Массив должен быть отсортирован по возрастанию
	sort.Strings(dictionary)

	//В результате каждое слово должно встречаться только один раз
	dictionary = getUnique(dictionary)

	//Все слова должны быть приведены к нижнему регистру
	for i, word := range dictionary {
		dictionary[i] = strings.ToLower(word)
	}

	for _, keyword := range dictionary {
		newAnagram := false
		s := make([]string, 0)
		for _, verifyWord := range dictionary {
			if isAnagram(keyword, verifyWord) {
				newAnagram = true
				s = append(s, verifyWord)
			}
		}

		//Множества из одного элемента не должны попасть в результат
		if newAnagram && len(s) > 1 {
			anagrams[s[0]] = s
		}
	}

	return anagrams
}

func main() {
	dictionary := []string{
		"Тяпка",
		"Пятак",
		"пятак",
		"пятка",
		"потоп",
		"топот",
		"топот",
		"слиток",
		"слиток",
		"столик",
		"листок",
		"пятак",
	}

	anagrams := searchAnagram(dictionary)

	for _, slice := range anagrams {
		for _, word := range slice {
			println(word)
		}
		println("")
	}
}
