# GoHelper &nbsp; ![GitHub release (latest by date)](https://img.shields.io/github/v/release/cipheras/gohelper?style=flat&logo=superuser)

#### A GO module to help in projects in generating formatted logs in log files and colored messages on the terminal. 

![Lines of code](https://img.shields.io/tokei/lines/github/cipheras/gohelper?style=flat)
&nbsp;&nbsp;&nbsp;&nbsp;![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cipheras/gohelper?style=flat)

## Installation
You can import this module and start using it.
```
go get github.com/cipheras/gohelper
```

## Usage
**gohelper** basically has two features, creating formatted logs in a file and show formatted texts on console.

### How to create formatted logs:
***Note:** You can configure `log` module according to your needs if you do not want to use gohelpers log format and still can use "Try" with it.*

You don't have to write `if err!=nil{}` everytime, you can just do
```
Try(error, mode, message)
where, mode can be true or false
If mode=true, process will exit and if mode=false, process will generate a warning message.
```
* This will write the same message to logs and also will show on terminal if some message is given.
* If no message is given, it will only log and do not show on terminal.
* To show original error on terminal, use `Cprint(err)`. If err!=nil, it will be shown on terminal.
* If you want to log `info`(other information), do:
```
Try(nil, false, message)
```

<br>To automatically generate `log.txt` file and write logs in it, call function `Flog()`.
```
Examples:
Try(err, true, "logging this") //it will log this message in case of err=nil, and if err!=nil then will log this message with error in file.
Try(err, false) //in this case it will not log in file if err=nil, and if err!=nil it will log the error.

OUTPUT:
2021/4/12 14:34:23 main.go:24: [INFO] generating logs
```

### How to create formatted texts on console:
* To show colors on **WINDOWS CMD** also, call function `Cwindows()`.
```
Example:
err := Cwindows()
```

```
Cprint(mode, message)
where mode can be,
    N = "normal"
    E = "error"
    W = "warning"
    T = "text"
    I = "info"
    S = "shell"
```
You can also use available colors and formats inbetween any of these or your own console outputs. 
<br>Available ones are:
```
RESET | RED | GREEN | YELLOW | BLUE | PURPLE | CYAN | WHITE | BGBLACK | BOLD | UNDERLINE | BLINK | CLEAR
Example:     
fmt.Println(BLUE, "hello", BOLD, var1, BLINK, var2, "!!", RESET)

Output:
```
![example](https://media.giphy.com/media/wSKT7XggpckVfTTkTB/giphy.gif)

## To Do
- [x] Colored text on console for linux
- [x] Colored text on cmd for windows
- [x] ERR,INFO,WARN type log structure
- [ ] Add more text and background colors

## License
**gohelper** is made by **@cipheras** and is released under the terms of the &nbsp;![GitHub License](https://img.shields.io/github/license/cipheras/gohelper?color=darkgreen)

## Contact &nbsp; [![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Fgohelper&label=Tweet)](https://twitter.com/intent/tweet?text=Hi:&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Fgohelper)

> Feel free to submit a bug, add features or issue a pull request.

