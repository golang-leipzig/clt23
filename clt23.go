package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	_ "embed"
)

//go:embed gopher.ascii
var gopher []byte

var message = `

██╗  ██╗███████╗██╗   ██╗     ██████╗██╗  ████████╗██████╗ ██████╗ ██╗
██║  ██║██╔════╝╚██╗ ██╔╝    ██╔════╝██║  ╚══██╔══╝╚════██╗╚════██╗██║
███████║█████╗   ╚████╔╝     ██║     ██║     ██║    █████╔╝ █████╔╝██║
██╔══██║██╔══╝    ╚██╔╝      ██║     ██║     ██║   ██╔═══╝  ╚═══██╗╚═╝
██║  ██║███████╗   ██║       ╚██████╗███████╗██║   ███████╗██████╔╝██╗
╚═╝  ╚═╝╚══════╝   ╚═╝        ╚═════╝╚══════╝╚═╝   ╚══════╝╚═════╝ ╚═╝

Welcome CLT 2023 visitor!

If you are interested in Go and cloud infrastructure topics, please visit us at
https://golangleipzig.space - our meetups are fun and well documented, it's not
far by 🚂 and we also regularly stream our events so you can join us via 📱 or
🖳, too.
`

func main() {
	fmt.Println(message)
	_, _ = io.Copy(os.Stdout, bytes.NewReader(gopher))
	fmt.Println("\nhttps://golangleipzig.space\n")
}
