package ui

import (
	"github.com/client/ui/chat"
	"github.com/client/ui/status"
	"github.com/client/util"
	"github.com/gotk3/gotk3/gtk"
)

/*
MainWindow is going to be the main window in which the chat application will display everything.
This window will consists of a @ChatView and a @StatusView, both inside a HBox, which isn't used for
anything else yet, so I won't add a refference to it on the struct itself

The Username of this MainWindow will be the username of the client.
*/
type MainWindow struct {
	*gtk.Window
	Chat   *chat.ChatView
	Status *status.StatusView
	Username string
}

func MainWindowNew(username string) *MainWindow {
	w, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	c, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)

	ch := chat.ChatViewNew()
	st := status.StatusViewNew()

	c.PackStart(st, false, false, uint(util.GetIntPortion(util.WIDTH, util.MARGIN)))
	c.PackEnd(ch, true, true, uint(util.GetIntPortion(util.WIDTH, util.MARGIN)))

	w.Add(c)
	w.SetSizeRequest(util.WIDTH, util.HEIGHT)

	w.Connect("delete-event", func() {
		gtk.MainQuit()
	})

	return &MainWindow{
		Window: w,
		Chat:   ch,
		Status: st,
		Username: username,
	}
}