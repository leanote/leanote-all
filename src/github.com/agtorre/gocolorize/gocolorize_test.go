package gocolorize

import (
	"fmt"
	"log"
	"testing"
)

var (
	DEBUG *log.Logger
	WARN  *log.Logger
)

func TestPaint(t *testing.T) {
	var blue Colorize

	//set some state
	blue.SetColor(Blue)

	out_string := fmt.Sprintf("Testing %s", blue.Paint("paint"))
	basis_string := "Testing \033[0;34mpaint\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestPaintString(t *testing.T) {
	var red Colorize

	//set some state
	red.SetColor(Red)

	out_string := red.Paint("Returning a string")
	basis_string := "\033[0;31mReturning a string\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestSetColorSetBgColor(t *testing.T) {
	var white_red_bg Colorize

	//set color and background
	white_red_bg.SetColor(White)
	white_red_bg.SetBgColor(Red)

	out_string := white_red_bg.Paint("Setting a foreground and background color!")
	basis_string := "\033[0;37m\033[41mSetting a foreground and background color!\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestPaintMultipleInterface(t *testing.T) {
	blue := Colorize{Fg: Blue}
	out_string := blue.Paint("Multiple types of args:", 1, 1.24)
	basis_string := "\033[0;34mMultiple types of args: 1 1.24\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestPaintComplexType(t *testing.T) {
	green := Colorize{Bg: Green}
	out_string := green.Paint("Complex types:",
		struct {
			int
			string
		}{})
	basis_string := fmt.Sprintf("\033[42mComplex types: %v\033[0m", struct {
		int
		string
	}{})
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestInitialize(t *testing.T) {
	black_on_white := Colorize{Fg: Black, Bg: White}
	f := black_on_white.Paint

	out_string := f("Now this is cool")
	basis_string := "\033[0;30m\033[47mNow this is cool\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestToggle(t *testing.T) {
	craziness := Colorize{Fg: Yellow, Bg: Black}
	craziness.ToggleFgIntensity()
	craziness.ToggleBgIntensity()
	craziness.ToggleBold()
	craziness.ToggleBlink()
	craziness.ToggleUnderline()
	craziness.ToggleInverse()

	out_string := craziness.Paint("craziness")
	basis_string := "\033[0;1;5;4;7;93m\033[100mcraziness\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}
func TestNewAllToggle(t *testing.T) {
	n := NewColor("yellow+bBuih:black+h")
	out_string := n.Paint("all toggles in 1 line!")
	basis_string := "\033[0;1;5;4;7;93m\033[100mall toggles in 1 line!\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestPlain(t *testing.T) {
	plain := Colorize{Fg: Magenta}
	SetPlain(true)
	out_string := plain.Paint("plain", "text")
	basis_string := "plain text"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
	SetPlain(false)
}
