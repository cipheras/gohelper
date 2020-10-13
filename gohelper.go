package gohelper

import (
	"fmt"
	"log"
	"os"
)

/*
RESET
RED
GREEN
YELLOW
BLUE
PURPLE
CYAN
WHITE
BGBLACK
BOLD
UNDERLINE
BLINK
CLEAR
LEFT
*/
const (
	// ANSI color codes
	RESET     = "\033[0m"
	RED       = "\033[38;5;1m"
	GREEN     = "\033[38;5;2m"
	YELLOW    = "\033[38;5;3m"
	BLUE      = "\033[38;5;4m"
	PURPLE    = "\033[38;5;5m"
	CYAN      = "\033[38;5;6m"
	WHITE     = "\033[38;5;7m"
	BGBLACK   = "\033[48;5;232m"
	BOLD      = "\033[1m"
	UNDERLINE = "\033[4m"
	BLINK     = "\033[5m"
	CLEAR     = "\033[2J\033[H"
	LEFT      = "\033[1000D"
)

const (
	// N :
	N = "normal"
	// E :
	E = "error"
	// W :
	W = "warning"
	// T :
	T = "text"
	// I :
	I = "info"
	// S :
	S = "shell"
)

// Try ...
func Try(err error, mode bool, msg ...interface{}) { //msgs,err,exit/noexit
	var msgs string
	for _, v := range msg {
		msgs = msgs + fmt.Sprintf("%v ", v)
	}
	if err != nil {
		if msgs != "" {
			if mode == true {
				Cprint(E, msgs)
				log.Println("[ERR]", msgs)
				os.Exit(0)
			}
			Cprint(W, msgs)
			log.Println("[WARN]", msgs)
			return
		}
		if mode == true {
			Cprint(E, err)
			log.Println("[ERR]", err)
			os.Exit(0)
		}
		Cprint(W, err)
		log.Println("[WARN]", err)
		return
	}
	if msgs != "" {
		log.Println("[INFO]", msgs)
	}
}

// Cprint ...
func Cprint(mode string, msg ...interface{}) {
	var msgs string
	for _, v := range msg {
		msgs = msgs + fmt.Sprintf("%v ", v)
	}
	switch mode {
	case N: //normal
		fmt.Println("\n" + CYAN + "[" + GREEN + "+" + CYAN + "] " + GREEN + msgs + RESET)
	case E: //error
		fmt.Println("\n" + CYAN + "[" + RED + BLINK + "-" + RESET + CYAN + "] " + RED + BGBLACK + BOLD + "ERROR" + RESET + " " + RED + msgs + RESET)
	case W: //warning
		fmt.Println("\n" + CYAN + "[" + YELLOW + BLINK + "!" + RESET + CYAN + "] " + YELLOW + BGBLACK + BOLD + "WARN" + RESET + " " + YELLOW + msgs + RESET)
	case T: //text
		fmt.Println("\n" + CYAN + "[" + PURPLE + "*" + CYAN + "] " + PURPLE + msgs + RESET)
	case I: //info
		fmt.Println("\n" + CYAN + "[" + BLUE + "i" + CYAN + "] " + BLUE + msgs + RESET)
	case S: //shell
		fmt.Print("\n" + CYAN + "[" + PURPLE + "*" + CYAN + "] " + PURPLE + msgs + "\n" + GREEN + ">> " + RESET)
	}
}

// Cwindows ...Edit registry to suppory ANSII
func Cwindows() error {
	return nil
}

/*
   Up: \u001b[{n}A moves cursor up by n
   Down: \u001b[{n}B moves cursor down by n
   Right: \u001b[{n}C moves cursor right by n
   Left: \u001b[{n}D moves cursor left by n

   Next Line: \u001b[{n}E moves cursor to beginning of line n lines down
   Prev Line: \u001b[{n}F moves cursor to beginning of line n lines down

   Set Column: \u001b[{n}G moves cursor to column n
   Set Position: \u001b[{n};{m}H moves cursor to row n column m

   Clear Screen: \u001b[{n}J clears the screen
       n=0 clears from cursor until end of screen,
       n=1 clears from cursor to beginning of screen
       n=2 clears entire screen
   Clear Line: \u001b[{n}K clears the current line
       n=0 clears from cursor to end of line
       n=1 clears from cursor to start of line
       n=2 clears entire line

   Save Position: \u001b[{s} saves the current cursor position
   Save Position: \u001b[{u} restores the cursor to the last saved position
*/
