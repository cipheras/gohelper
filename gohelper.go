package gohelper

import (
	"fmt"
	"log"
	"os"
)

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

func Flog() error {
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(f)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	// defer f.Close()
	return nil
}

// Try ...
func Try(err error, mode bool, msg ...interface{}) { //err,exit/noexit,msg
	var msgs string
	for _, v := range msg {
		msgs = msgs + fmt.Sprintf("%v ", v)
	}
	if err != nil {
		if msgs != "" {
			if mode == true {
				Cprint(E, msgs)
				log.Println("[ERR]", msgs, err)
				os.Exit(0)
			}
			Cprint(W, msgs)
			log.Println("[WARN]", msgs, err)
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
	} else if msgs != "" {
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

const (
	Up    = "\u001b[{n}A" //moves cursor up by n
	Down  = "\u001b[{n}B" //moves cursor down by n
	Right = "\u001b[{n}C" //moves cursor right by n
	Left  = "\u001b[{n}D" //moves cursor left by n

	NextLine = "\u001b[{n}E" //moves cursor to beginning of line n lines down
	PrevLine = "\u001b[{n}F" //moves cursor to beginning of line n lines down

	SetColumn   = "\u001b[{n}G"     //moves cursor to column n
	SetPosition = "\u001b[{n};{m}H" //moves cursor to row n column m

	ClearScreen = "\u001b[{n}J"
	// clears the screen
	//     n=0 clears from cursor until end of screen,
	//     n=1 clears from cursor to beginning of screen
	//     n=2 clears entire screen

	ClearLine = "\u001b[{n}K"
	// clears the current line
	//     n=0 clears from cursor to end of line
	//     n=1 clears from cursor to start of line
	//     n=2 clears entire line

	SavePosition    = "\u001b[s" //saves the current cursor position
	RestorePosition = "\u001b[u" //restores the cursor to the last saved position
)
