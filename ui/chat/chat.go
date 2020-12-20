package chat

import (
	"github.com/client/util"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

/*
ChatMessageArea will be the element that takes charge of displaying the
area in which messages are seen, sent and responded. This area will be composed
of a frame, from which derives a scrolled window with a VBox as it's child. This
VBox will then hold @ChatMessage objects, that will be shown.

The MESSAGE_SPACING Variable will put the spacing in between the messages on the VBox.
The BORDER_SPACING Variable will put the spacing in between the messages and the borders of the VBox
The MAX_MESSAGES variable will have how many elements the VBox has until it starts removing messages to load new ones.
*/
type ChatMessageArea struct {
	*gtk.Frame
	Scroll *gtk.ScrolledWindow
	Box    *gtk.Box
}

const (
	MESSAGE_SPACING int  = 10
	BORDER_SPACING  uint = 10
	MAX_MESSAGES    uint = 30
)

var VIEW_WIDTH int = util.GetIntPortion(util.WIDTH, 0.7)

func ChatMessageAreaNew() *ChatMessageArea {
	f, _ := gtk.FrameNew("Chat Window")
	s, _ := gtk.ScrolledWindowNew(nil, nil)
	s.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_ALWAYS)
	b, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, MESSAGE_SPACING)

	b.SetVAlign(gtk.ALIGN_END)
	f.SetSizeRequest(VIEW_WIDTH, util.GetIntPortion(util.HEIGHT, 0.7))
	f.Add(s)
	s.Add(b)
	return &ChatMessageArea{
		Frame:  f,
		Scroll: s,
		Box:    b}
}

//Displays the message through the ChatMessageArea
func (chat *ChatMessageArea) SendMessage(owner string, message string) error {
	l := chat.Box.GetChildren().Length()
	if l >= MAX_MESSAGES {
		_, err := glib.IdleAdd(func() {
			chat.Box.GetChildren().Nth(0).Data().(*gtk.Widget).Destroy()
		})
		if err != nil {
			return err
		}
	}

	_, err := glib.IdleAdd(func() {
		message := NewChatMessage(owner, message)
		chat.Box.PackStart(message, false, false, uint(BORDER_SPACING))
		message.ShowAll()
	})

	if err != nil {
		return err
	}
	return nil
}