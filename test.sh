#!/bin/bash
echo "生成golang 1.3版本的pot文件..."
./godoc-i18n go1.3/io.go
echo "golang 1.3版本的pot文件已经生成：go1.3/io.pot"
#翻译
echo "翻译golang 1.3版本的po文件中..."
echo "golang 1.3版本的po文件已经翻译：io.po"
#升级golang 1.4版本
echo
echo "golan 1.4版本发布了"
echo "生成golang 1.4版本的pot文件..."
./godoc-i18n go1.4/io.go
echo "golang 1.4版本的pot文件已经生成：go1.4/io.pot"
echo "更新golang 1.3版本的翻译成果到1.4版本中..."
msgmerge io.po go1.4/io.pot > io-1.4.po
echo "golang 1.4版本po文件已经生成，并集成了1.3版本的翻译成果：io-1.4.po"
