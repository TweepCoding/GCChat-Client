package chat

import (
	"github.com/gotk3/gotk3/gtk"
)

/*
ChatView will be the main view for the area in which the chat takes place. The chat view won't have a preffered size,
and instead, will just fill the area that hasn't been occupied by the @StatusView. The ChatView is composed of a box
containing a @ChatMessageArea and a @ChatInput. The @ChatMessageArea will 
*/
type ChatView struct {
	*gtk.Box
	Area *ChatMessageArea
	Input *ChatInput
}

const (
	VIEW_SPACING int = 20
	VIEW_PADDING uint = 20
)

func ChatViewNew() *ChatView {
	c, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, VIEW_SPACING)
	ch := ChatMessageAreaNew()
	in := ChatInputNew()
	c.PackStart(ch, true, true, VIEW_PADDING)
	c.PackEnd(in, false, false, VIEW_PADDING)
	return &ChatView{
		Box: c,
		Area: ch,
		Input: in,
	}
}
