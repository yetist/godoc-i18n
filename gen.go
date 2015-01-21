package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"strings"
)

func genEntry(filename, used, text string, start, end int) {
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
	fmt.Printf(format, used, filename, start, lens, msgid)
}

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "/usr/lib/go/src/io/io.go", nil, parser.ParseComments)
	if err != nil {
		return
	}
	fmt.Println(f.Decls)
	for _, s := range f.Comments {
		start_pos := fset.Position(s.Pos())
		end_pos := fset.Position(s.End())
		genEntry("io.go", "func Hi", s.Text(), start_pos.Line, end_pos.Line)
	}
}
