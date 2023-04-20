package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Args struct {
	k int
	n bool
	r bool
	u bool
}

type flagN struct {
	num int
	str string
}

func getFlags() (*Args, error) {
	k := flag.Int("k", 0, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	args := &Args{
		k: *k,
		n: *n,
		r: *r,
		u: *u,
	}

	if args.k < 0 {
		return nil, errors.New("number cannot be negative")
	}

	return args, nil
}

// записывает строки из файла в слайс
func readLines(file string) (lines []string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		const delim = '\n'
		line, err := r.ReadString(delim)
		if err == nil || len(line) > 0 {
			if err != nil {
				line += string(delim)
			}
			lines = append(lines, line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return lines, nil
}

// сортирует слайс
func Sort(lines []string) ([]string, error) {
	args, err := getFlags()
	if err != nil {
		return []string{}, err
	}

	//flag -k
	swap(lines, args.k)

	//flag -n
	if args.n {
		lines = sortInt(lines)
	} else {
		sort.Strings(lines)
	}
	//flag -k
	swap(lines, args.k)
	//flag -u
	if args.u {
		lines = getUnique(lines)
	}
	//flag -r
	if args.r {
		lines = reverseStrs(lines)
	}
	return lines, nil
}

// сортировка по числовому значению
func sortInt(slice []string) []string {
	const zeroUnicode, nineUnicode = 48, 57

	nums := make([]flagN, 0, len(slice))
	for _, s := range slice {
		n := flagN{str: s}
		runes := []rune(s)
		for i, r := range runes {
			if r < zeroUnicode || r > nineUnicode {
				n.num, _ = strconv.Atoi(s[:i])
				n.str = s
				break
			}
		}
		nums = append(nums, n)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i].num < nums[j].num
	})
	for i := 0; i < len(slice); i++ {
		slice[i] = nums[i].str
	}
	return slice
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

// возвращает слайс наоборот
func reverseStrs(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseStrs(input[1:]), input[0])
}

// меняет местами lines[0] и lines[numWord]
func swap(lines []string, numWord int) {
	for i, line := range lines {
		var newLine string
		strs := strings.Split(line, " ")
		strs[0], strs[numWord] = strs[numWord], strs[0]
		newLine = strs[0]
		for i := 1; i < len(strs); i++ {
			newLine += " " + strs[i]
		}

		lines[i] = newLine
	}
}

// пишет каждый элемент слайса в файл
func writeLines(file string, lines []string) (err error) {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, line := range lines {
		_, err := w.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	file := `lines.txt`
	lines, err := readLines(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines, err = Sort(lines)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = writeLines(file, lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
