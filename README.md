# godoc-i18n

提供golang doc翻译方案

请运行./test.sh 模拟翻译过程

模拟翻译过程如下：

```
生成golang 1.3版本的pot文件...
golang 1.3版本的pot文件已经生成：go1.3/io.pot

翻译golang 1.3版本的po文件中...
#这里可使用poedit，transifex平台等进行翻译
golang 1.3版本的po文件已经翻译：io.po

golan 1.4版本发布了

生成golang 1.4版本的pot文件...
golang 1.4版本的pot文件已经生成：go1.4/io.pot

更新golang 1.3版本的翻译成果到1.4版本中...
golang 1.4版本po文件已经生成，并集成了1.3版本的翻译成果：io-1.4.po
```

翻译完成之后，请对比io.po和io-1.4.po，分别对应于golang 1.3和1.4版本，这是翻译成果。

后续方案设想：

* 根据po文件中的行号信息，直接将翻译的中文替换到源文件中，如下面示例中的#: go1.4/io.go:480:5，表示io.go文件中从第480行开始的5行内容需要替换掉。
```
#. used by: func()这是上下文:
#: go1.4/io.go:480:5
msgid ""
"TeeReader returns a Reader that writes to w what it reads from r.All reads "
"from r performed through it are matched withcorresponding writes to w.  "
"There is no internal buffering -the write must complete before the read "
"completes.Any error encountered while writing is reported as a read error."
msgstr "翻译测试"
```

对于系统中本身已经存在了golang，源代码是英文的用户，可将po文件和相应的合并工具单独发布，由用户自行进行合并工具将翻译成果合并到其源代码中。

* 不将翻译结果合并到源文件中，而是修改godoc，使其能支持搜索存在于特定目录中的po文件，在运行godoc时，从po文件中提取中文进行显示。
