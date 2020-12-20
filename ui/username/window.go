package username

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/client/ui"
	"github.com/client/conn"
)

/*
UsernameWindow isn't a type, but regardless, it's a window that will ask the user for input
about their name and the server they wish to connect to. The function will simply prebuild the
window.

UWIDTH gives the width of the username window
UHEIGHT gives the height of the username window
UMARGIN gives the margin of the username window
*/

const (
	UWIDTH  int  = 500
	UHEIGHT int  = 200
	UMARGIN uint = 10
)

func UsernameWindowNew() *gtk.Window {
	w, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	c, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, int(UMARGIN))
	b, _ := gtk.ButtonNew()
	e, _ := gtk.EntryNew()
	l, _ := gtk.LabelNew("Enter username:")

	w.Connect("delete-event", func() {
		gtk.MainQuit()
	})

	b.Connect("clicked", func() {
		w.Destroy()
		t, _ := e.GetText()
		m := ui.MainWindowNew(t)
		conn.Connect(m)
		m.ShowAll()
	})

	b.SetLabel("Start")

	c.PackStart(l, true, true, UMARGIN)
	c.PackStart(e, true, true, UMARGIN)
	c.PackStart(b, true, true, UMARGIN)

	c.SetMarginStart(int(UMARGIN))
	c.SetMarginEnd(int(UMARGIN))

	w.Add(c)
	w.SetResizable(false)
	w.SetPosition(gtk.WIN_POS_CENTER)
	w.SetDefaultSize(UWIDTH, UHEIGHT)
	w.SetTitle("Username Entry")

	return w
}
