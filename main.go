package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	var (
		hashtype = flag.String("type", "plain", "Hash Type (plain(nohash), md5, sha1, sha256, sha512)")
	)
	flag.Parse()

	var stdin string
	fmt.Scan(&stdin)

	p := []byte(stdin)

	var hashed_string string
	if *hashtype == "plain" {
		hashed_string = stdin
	} else if *hashtype == "md5" {
		hashed_string = fmt.Sprintf("%x", md5.Sum(p))
	} else if *hashtype == "sha1" {
		hashed_string = fmt.Sprintf("%x", sha1.Sum(p))
	} else if *hashtype == "sha256" {
		hashed_string = fmt.Sprintf("%x", sha256.Sum256(p))
	} else if *hashtype == "sha512" {
		hashed_string = fmt.Sprintf("%x", sha512.Sum512(p))
	} else {
		hashed_string = stdin
	}

	for _, c := range hashed_string {
		moji(string(c))
	}
}

func moji(c string) {
	/*
		cred := color.New(color.Reset, color.BgHiRed)
		cblue := color.New(color.Reset, color.BgHiBlue)
		ccyan := color.New(color.Reset, color.BgHiCyan)
		cgreen := color.New(color.Reset, color.BgHiGreen)
		cmagenta := color.New(color.Reset, color.BgHiMagenta)
		cwhite := color.New(color.Reset, color.BgHiWhite)
		cyellow := color.New(color.Reset, color.BgHiYellow)
		cblack := color.New(color.Reset, color.BgHiBlack)
	*/
	colour := color.New(color.Reset)
	if c == "0" {
		colour = color.New(color.Reset, color.BgHiRed)
	} else if c == "1" {
		colour = color.New(color.Reset, color.BgHiBlue)
	} else if c == "2" {
		colour = color.New(color.FgBlack, color.BgHiCyan)
	} else if c == "3" {
		colour = color.New(color.FgBlack, color.BgHiGreen)
	} else if c == "4" {
		colour = color.New(color.Reset, color.BgHiMagenta)
	} else if c == "5" {
		colour = color.New(color.FgBlack, color.BgHiWhite)
	} else if c == "6" {
		colour = color.New(color.FgBlack, color.BgHiYellow)
	} else if c == "7" {
		colour = color.New(color.Reset, color.BgHiBlack)
	} else if c == "8" {
		colour = color.New(color.Reset, color.BgRed)
	} else if c == "9" {
		colour = color.New(color.Reset, color.BgBlue)
	} else if c == "a" || c == "A" {
		colour = color.New(color.Reset, color.BgCyan)
	} else if c == "b" || c == "B"{
		colour = color.New(color.FgBlack, color.BgGreen)
	} else if c == "c" || c == "C"{
		colour = color.New(color.Reset, color.BgMagenta)
	} else if c == "d" || c == "D"{
		colour = color.New(color.FgBlack, color.BgWhite)
	} else if c == "e" || c == "E"{
		colour = color.New(color.FgBlack, color.BgYellow)
	} else if c == "f" || c == "F"{
		colour = color.New(color.FgWhite, color.BgBlack)
	}
	colour.Print(c)
}
