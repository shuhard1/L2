package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
)

// wordCache — это кеш в памяти нашего списка слов
var wordCache []string

type match struct {
	Before string
	Match  string
	After  string
}

type Args struct {
	A int
	B int
	C int
	c bool
	v bool
	F bool
	n bool
}

func getFlags() (*Args, error) {
	A := flag.Int("A", 0, "'after' печатать +N строк после совпадения")
	B := flag.Int("B", 0, "'before' печатать +N строк до совпадения")
	C := flag.Int("C", 0, "'context' (A+B) печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "'count' (количество строк)") //len(match)
	v := flag.Bool("v", false, "'invert' (вместо совпадения, исключать)")
	F := flag.Bool("F", false, "'fixed', точное совпадение со строкой, не паттерн") //тупое сравнение строк ==
	n := flag.Bool("n", false, "'line num', напечатать номер строки")               //в цикле for i := 0; i < len(words); i++ { создай счетик там где совпадение запоминай счетчик
	flag.Parse()

	args := &Args{
		A: *A,
		B: *B,
		C: *C,
		c: *c,
		v: *v,
		F: *F,
		n: *n,
	}

	if args.A < 0 || args.B < 0 || args.C < 0 {
		return nil, errors.New("number cannot be negative")
	}

	return args, nil
}

// readLines читает весь файл в память
// и возвращает часть своих строк.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// grepDictionary фильтрует слова по шаблону.
// Возвращает ошибку, если выражение неверно.
func grepDictionary(pattern string, words []string) ([]match, int, []int, error) {
	args, err := getFlags()
	if err != nil {
		return nil, 0, nil, err
	}
	//обьект Regexp можно использовать для сопоставления с текстом
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, 0, nil, err
	}
	var matches []match
	var B, A int = args.B, args.A
	if args.C > 0 {
		B, A = args.C, args.C
	}

	var counter bool
	if args.c {
		counter = true
	}
	var count int

	var invert bool
	if args.v {
		invert = true
	}

	var fixed bool
	if args.F {
		fixed = true
	}

	var lineCounter bool
	if args.n {
		lineCounter = true
	}
	var lineNum []int
	for i := 0; i < len(words); i++ {
		var N int = -B
		if loc := re.FindStringIndex(words[i]); loc != nil && !invert || invert && loc == nil {
			if fixed && pattern != words[i] {
				continue
			}
			if invert {
				m := match{
					Before: "",
					Match:  words[i],
					After:  ""}
				matches = append(matches, m)
				continue
			}
			for ; N <= A; N++ {
				//проверка на index out of range
				if i+N >= len(words) || i+N < 0 {
					continue
				}

				if N == 0 && counter {
					count++
				}

				if N == 0 && lineCounter {
					lineNum = append(lineNum, i)
				}

				m := match{
					Before: words[i+N][:loc[0]],
					Match:  words[i+N][loc[0]:loc[1]],
					After:  words[i+N][loc[1]:]}
				matches = append(matches, m)
			}
		}
	}
	return matches, count, lineNum, nil
}

// getWords возвращает кэшированный список слов или загружает его.
func getWords(file string) ([]string, error) {
	if wordCache == nil {
		words, err := readLines(file)
		if err != nil {
			return nil, err
		}
		wordCache = words
	}
	return wordCache, nil
}

func Gorp(word string, file string) {
	words, err := getWords(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result, count, lineNum, err := grepDictionary(word, words)
	if err != nil {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if count > 0 {
		println(count)
		return
	}
	if lineNum != nil {
		for _, num := range lineNum {
			println(num + 1)
		}
		return
	}
	for _, math := range result {
		println(math.Before + math.Match + math.After)
	}
}

func main() {
	Gorp("Please", "test1.txt")
}
