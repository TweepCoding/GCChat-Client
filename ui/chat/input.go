package chat

import (
	"github.com/gotk3/gotk3/gtk"
)

/*
ChatInput will be the object that manages the entry of chat messages, and how
they are going to be sent. The sent messages will be seen globally by every
user through the @ChatMessageArea that is part of the @ChatView of each user
*/
type ChatInput struct {
	*gtk.Box
	Entry *gtk.Entry
}

const ENTRY_SPACING int = 0

func ChatInputNew() *ChatInput {
	e, _ := gtk.EntryNew()
	c, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, ENTRY_SPACING)
	c.PackStart(e, true, true, 0)
	return &ChatInput{
		Box: c,
		Entry: e,
	}
}