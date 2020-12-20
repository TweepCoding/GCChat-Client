package chat

import (
	"github.com/client/util"
	"github.com/gotk3/gotk3/gtk"
)

/*
ChatMessage will be the element that is inside @ChatMessageArea and will be the one shown
to the user when it recieves a message. This element will be a box that contains the owner of
the message sent, as well as the text that was sent by the aformentioned owner of the message.

The ELEMENT_SPACING variable will be the spacing in between the message and the owner's username
*/
type ChatMessage struct {
	*gtk.Box
	Owner string
	Text string
}

const ELEMENT_SPACING int = 10

func NewChatMessage(owner string, message string) *ChatMessage {
	b, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, ELEMENT_SPACING)
	o, _ := gtk.LabelNew(owner + ":")
	l, _ := gtk.LabelNew(message)
	l.SetHAlign(gtk.ALIGN_START)
	l.SetLineWrap(true)
	b.SetSizeRequest(VIEW_WIDTH, util.GetIntPortion(util.HEIGHT, 0.05))
	b.PackStart(o, false, false, 0)
	b.PackEnd(l, true, true, 0)
	return &ChatMessage{
		Box: b,
		Owner: owner,
		Text: message,
	}
}