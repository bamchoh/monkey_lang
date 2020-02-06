package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bamchoh/monkey_lang/evaluator"
	"github.com/bamchoh/monkey_lang/lexer"
	"github.com/bamchoh/monkey_lang/object"
	"github.com/bamchoh/monkey_lang/parser"
)

func load(filename string) (string, error) {
	fullpath, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}
	fp, err := os.Open(fullpath)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	script, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", err
	}

	return string(script), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify monkey lang script file")
		os.Exit(-1)
	}

	script, err := load(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	out := os.Stdout
	env := object.NewEnvironment()
	l := lexer.New(script)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		os.Exit(-1)
	}
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil && evaluated.Type() == object.ERROR_OBJ {
		io.WriteString(out, " evaluator errors:\n")
		fmt.Println("\t" + evaluated.Inspect())
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
