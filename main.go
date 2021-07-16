package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type TemplateContext struct {
	Args map[string]string
}

func main() {
	if len(os.Args) < 2 {
		panic("need a filename")
	}
	filename := os.Args[1]
	args := make(map[string]string)
	if len(os.Args) >= 3 {
		_args := os.Args[2:]
		for _, a := range _args {
			splitted := strings.Split(a, "=")
			args[splitted[0]] = splitted[1]
		}
	}
	t := template.New("template")
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	t, err = t.Parse(string(bs))
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, &TemplateContext{Args: args})
	if err != nil {
		panic(err)
	}

}
