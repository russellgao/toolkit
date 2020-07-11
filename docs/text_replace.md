# text_replace 文本替换

## 用法
```
tkctl text_replace --help
文本替换，支持正则替换和非正则替换，类似与linux下的sed，但比sed更好用，而且可以跨平台使用

Usage:
  tkctl text_replace [flags]

Flags:
  -h, --help                  help for text_replace
      --text-dirs string      需要替换的目录, 默认为当前路径 (default ".")
      --text-mode string      替换的模式，支持正则（regexp）和非正则（text）两种模式，默认非正则， (default "text")
      --text-pattern string   需要替换的pattern  [required]
      --text-repl string      目标字符串  [required]
```

## 例子
### 把docs目录下的 所有 `golang` 替换成 `python`
```
tkctl text_replace --text-dirs docs --text-pattern golang --text-repl python
```

### 把docs目录下的 所有 `g.*g` 正则替换成 `python`
```
tkctl text_replace --text-dirs docs --text-pattern g.*g --text-repl python --text-mode regexp
```
