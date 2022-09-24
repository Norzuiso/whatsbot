package main

import qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"

func qr() {
	qr := make(chan string)
	obj := qrcodeTerminal.New()
	obj.Get(qr).Print()
}
