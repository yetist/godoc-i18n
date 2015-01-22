package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func genEntry(file *os.File, filename, used, text string, start, end int) {
	format := `
#. used by: %s:
#: %s:%d:%d
msgid ""
"%s"
msgstr ""
`
	lens := end - start + 1
	strs := strings.Split(text, "\n")
	m := strs[:len(strs)-1]
	for k, v := range m {
		if len(v) == 0 {
			m[k] = "\\n"
		}
	}
	msgid := strings.Join(m, "\"\n\"")
	content := fmt.Sprintf(format, used, filename, start, lens, msgid)
	file.Write([]byte(content))
}

var header = `# Chinese translations for PACKAGE package
# PACKAGE 软件包的简体中文翻译.
# Copyright (C) 2015 THE PACKAGE'S COPYRIGHT HOLDER
# This file is distributed under the same license as the PACKAGE package.
# yetist <yetist@gmail.com>, 2015.
#
msgid ""
msgstr ""
"Project-Id-Version: PACKAGE VERSION\n"
"Report-Msgid-Bugs-To: \n"
"POT-Creation-Date: 2015-01-11 17:06+0800\n"
"PO-Revision-Date: 2015-01-11 16:25+0800\n"
"Last-Translator: yetist <yetist@gmail.com>\n"
"Language-Team: Chinese (simplified)\n"
"Language: zh_CN\n"
"MIME-Version: 1.0\n"
"Content-Type: text/plain; charset=UTF-8\n"
"Content-Transfer-Encoding: 8bit\n"
`

func main() {
	filename := os.Args[1]
	var file, out *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	outfile := strings.TrimSuffix(filename, "go") + "pot"
	//outfile := strings.Replace(filename, "go", "pot", -1)
	if out, err = os.Create(outfile); err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return
	}
	out.Write([]byte(header))
	for _, s := range f.Comments {
		start_pos := fset.Position(s.Pos())
		end_pos := fset.Position(s.End())
		genEntry(out, filename, "func()这是上下文", s.Text(), start_pos.Line, end_pos.Line)
	}
}
