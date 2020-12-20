package main

import (
	"github.com/client/ui/username"
	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(nil)
	win := username.UsernameWindowNew()

	win.ShowAll()
	gtk.Main()
}
