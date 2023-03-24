package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Args struct {
	f string
	d string
	s bool

	files []string
}

func getArgs() (*Args, error) {
	f := flag.String("f", "", "'fields' - выбрать поля (колонки)")
	d := flag.String("d", "\t", "'delimiter' - использовать другой разделитель")
	s := flag.Bool("s", false, "'separated' - только строки с разделителем")

	flag.Parse()

	args := &Args{
		f: *f,
		d: *d,
		s: *s,
	}

	args.files = append(args.files, flag.Args()...)

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
func getF(args Args) ([]int, error) {
	var result []int

	for i := 0; i < len(strings.Split(args.f, ",")); i++ {
		num, err := strconv.Atoi(strings.Split(args.f, ",")[i])
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}

func cut() ([]string, error) {

	var result []string

	args, err := getArgs()
	if err != nil {
		return nil, err
	}

	f, err := getF(*args)
	if err != nil {
		return nil, err
	}

	for _, fileName := range args.files {
		lines, err := readLines(fileName)
		if err != nil {
			return nil, err
		}

		var r []string

		for _, line := range lines {
			if args.d != "" && strings.Contains(line, args.d) {
				words := strings.Split(line, args.d)

				cutLine := strings.Builder{}

				for _, val := range f {
					if len(words) >= val {
						cutLine.WriteString(words[val-1])
						cutLine.WriteString(args.d)
					}
				}

				r = append(r, strings.TrimSuffix(cutLine.String(), args.d))

			} else if !args.s {
				r = append(r, line)
			}
		}
		cutLines := r
		result = append(result, cutLines...)
	}

	return result, nil
}

func main() {
	lines, err := cut()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := range lines {
		fmt.Println(lines[i])
	}
}
