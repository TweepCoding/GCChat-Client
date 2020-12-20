package status

import (
	"github.com/gotk3/gotk3/gtk"
)

/*
StatusMember is going to be the elements that will go inside the @StatusView
of the @MainWindow of the program. For now, they will only contain the user that
has logged on, but in the future, they will contain more elements
*/
type StatusMember struct {
	*gtk.Box
	Name string
}

const (
	MEMBER_SPACING int = 0
	MEMBER_HEIGHT int = 30
)

func StatusMemberNew(name string) *StatusMember {
	c, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, MEMBER_SPACING)
	l, _ := gtk.LabelNew(name)
	c.PackStart(l, true, true, 0)
	c.SetSizeRequest(VIEW_WIDTH, MEMBER_HEIGHT)
	return &StatusMember{
		Box:  c,
		Name: name,
	}
}
