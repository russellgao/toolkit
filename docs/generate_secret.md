# generate_secret 生成随机密码

## 用法
```
 tkctl generate_secret --help 
生成随机密码，支持1～64为长度，可以指定是否包含特殊字符

Usage:
  tkctl generate_secret [flags]

Flags:
  -h, --help                  help for generate_secret
      --length int            随机密码的长度 (default 16)
      --special_char string   是否包含特殊字符，默认为包含所有的特殊字符。如果不包含请传值为false，此时随机密码由数字和大小写字母组成；如果需要自定义特殊请通过此参数直接传入即可，例:只包含 #*& 的特殊字符，此参数应该为 --special_char #*&

```

## 例子
### 生成包含特殊字符的16为随机密码

```
tkctl generate_secret
```

### 生成包含特殊字符的20为随机密码
```
tkctl generate_secret --length 20
```

### 生成不包含特殊字符的20为随机密码
```
tkctl generate_secret --length 20 --special_char false 
```

### 生成仅包含 `*&^%$` 特殊字符的 32 为随机密码
```
tkctl generate_secret --length 20 --special_char '*&^%$'
```