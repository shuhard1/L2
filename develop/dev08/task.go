package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

// Cd меняет директорию
func Cd(argv []string) error {
	if len(argv) == 1 {
		os.Chdir(os.Getenv("HOME"))
	} else if argv[1][0:1] == "/" {
		os.Chdir(argv[1])
	} else if argv[1][0:1] == "~" {
		os.Chdir(os.Getenv("HOME") + strings.Join(argv[1:], ""))
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		os.Chdir(wd + "/" + argv[1])
	}
	return nil
}

// Exit выход из shell
func Exit(argv []string) error {
	if len(argv) == 1 {
		os.Exit(0)
	} else if val, err := strconv.Atoi(argv[1]); err == nil {
		os.Exit(val)
	} else {
		return errors.New("exit: exit codes are integers")
	}
	return nil
}

// Check проверяет, является ли команда встроенной командой
func Check(argv []string) (func([]string) error, error) {
	switch argv[0] {
	case "cd":
		return Cd, nil
	case "exit":
		return Exit, nil
	default:
		return nil, errors.New(argv[0] + " не является встроенной командой")
	}
}

type BuiltinPromptItem struct {
	Username      string
	Hostname      string
	Commandprompt string
}

func (i BuiltinPromptItem) Generate() Item {
	user, _ := user.Current()
	i.Username = user.Username
	i.Hostname, _ = os.Hostname()
	i.Commandprompt = ": "
	return i
}

func (i BuiltinPromptItem) String() string {
	wd, _ := os.Getwd()
	if wd == os.Getenv("HOME") {
		wd = "~"
	} else if strings.HasPrefix(wd, os.Getenv("HOME")) {
		wd = "~" + strings.TrimPrefix(wd, os.Getenv("HOME"))
	}
	return strings.Join([]string{
		i.Username,
		"@",
		i.Hostname,
		" ",
		wd,
		i.Commandprompt,
	}, "")
}

func (i BuiltinPromptItem) Prefix() {}

func Promptt() Prompt {
	return Prompt{
		Items: []Item{
			BuiltinPromptItem{},
		},
	}
}

// Item представляет, как должен вести себя элемент подсказки
type Item interface {
	Generate() Item
	Prefix()
	String() string
}

// Prompt содержит то, что необходимо отобразить, прежде чем запрашивать ввод данных пользователем.
type Prompt struct {
	Items []Item
}

// Print печатает приглашение перед ожиданием ввода пользователя
func (p *Prompt) Print() {
	for i := 0; i < len(p.Items); i++ {
		if i != 0 {
			p.Items[i].Prefix()
		}
		fmt.Print(p.Items[i])
	}
}

// Generate инициализирует все элементы подсказки
func (p *Prompt) Generate() {
	for i := 0; i < len(p.Items); i++ {
		p.Items[i] = p.Items[i].Generate()
	}
}

const (
	// shellLoginText содержит текст, который выводится при инициализации оболочки.
	shellLoginText = "Gosh v0.1.0\n"
)

func initialize() {
	fmt.Print(shellLoginText)

	// todo: добавить конфигурацию

	// инициализировать reader
	reader = bufio.NewReader(os.Stdin)

	// загрузить подсказку с элементами подсказки
	// todo: загрузить элементы подсказки, если подсказка настроена в файле конфигурации
	shellPrompt = Promptt()

	// выводит содержимое каждого элемента подсказки
	shellPrompt.Generate()
}

func interpret() error {
	// todo: добавить прослушиватель сигнала

	// распечатать приглашение
	shellPrompt.Print()

	// читать пользовательский ввод
	input, err := readInput()

	// return errors
	if err != nil {
		if err == io.EOF {
			return err
		}
		return err
	}

	input = strings.TrimRight(input, "\r\n")

	// если вход не указан, пропустить цикл
	if input == "" {
		return nil
	}

	// сделать: добавить историю

	// разделить ввод в аргументах
	argv := strings.Fields(input)

	// проверьте, является ли команда встроенной командой
	fn, err := Check(argv)
	if err == nil {
		err = fn(argv)
		return err
	}

	// в противном случае выполните команду
	cmd := exec.Command("bash", "-c", input)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		return err
	}

	return nil

}

func readInput() (string, error) {
	return reader.ReadString('\n')
}

var (
	// shellPrompt содержит то, что отображается каждый раз, когда пользователю предлагается ввести данные.
	shellPrompt Prompt
	// reader reads the user input
	reader *bufio.Reader
)

// Run содержит весь жизненный цикл оболочки
func Run() {

	initialize()

	// бесконечно интерпретировать каждый ввод
	for {
		err := interpret()
		if err == nil {
			continue
		}
		if err == io.EOF {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		fmt.Println(err)
	}

}

func main() {
	Run()
}
