package main

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed go
var f embed.FS

type TemplateContext struct {
	Args map[string]string
}

func getTemplateContent(name string) (string, error) {
	u, err := url.Parse(name)
	if err != nil {
		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil {
			return "", err
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(bs), nil
	} else {
		fmt.Println("in file mode")
		path, err := filepath.Abs(name)
		if err != nil {
			return "", err
		}
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("in internal template mode")
			bs, err := f.ReadFile(name)
			if err != nil {
				return "", err
			}
			return string(bs), nil
		}
		bs, err := ioutil.ReadFile(path)
		if err != nil {
			return "", err
		}
		return string(bs), nil
	}
}

func Make(filename string, args map[string]string) (string, error) {
	content, err := getTemplateContent(filename)
	if err != nil {
		return "", err
	}
	t := template.New("template")

	t, err = t.Parse(content)
	if err != nil {
		return "", err
	}

	b := bytes.NewBuffer([]byte{})
	err = t.Execute(b, &TemplateContext{Args: args})
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func main() {
	if len(os.Args) < 2 {
		panic("need a template filename")
	}
	args := make(map[string]string)
	if len(os.Args) >= 3 {
		_args := os.Args[2:]
		for _, a := range _args {
			splitted := strings.Split(a, "=")
			args[splitted[0]] = splitted[1]
		}
	}
	o, err := Make(os.Args[1], args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(o)
}
