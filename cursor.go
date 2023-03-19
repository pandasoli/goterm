package goterm
import "fmt"


func HideCursor() {
  fmt.Print("\033[?25l")
}

func ShowCursor() {
  fmt.Print("\033[?25h")
}

func IBeam_cursor() { fmt.Print("\033[6 q") }
func Block_cursor() { fmt.Print("\033[2 q") }
func Underline_cursor() { fmt.Print("\033[4 q") }

func Blinking_IBeam_cursor() { fmt.Print("\033[5 q") }
func Blinking_Block_cursor() { fmt.Print("\033[1 q") }
func Blinking_Underline_cursor() { fmt.Print("\033[3 q") }
