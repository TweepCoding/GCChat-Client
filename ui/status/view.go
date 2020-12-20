package status

import (
	"sync"
	"github.com/gotk3/gotk3/glib"
	"github.com/client/util"
	"github.com/gotk3/gotk3/gtk"
)

/*
StatusView will be the view used for the status and notif area. This will
handle to see if people are online or not. In the future, this may be changed
to add more features
*/
type StatusView struct {
	*gtk.Frame
	Box *gtk.Box
	Users []*StatusMember
	Mutex *sync.RWMutex
}

const (
	VIEW_SPACING int = 0
	VIEW_WIDTH int = 200
)

//Creates a new StatusView
func StatusViewNew() *StatusView {
	f, _ := gtk.FrameNew("Online:")
	c, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, VIEW_SPACING)
	f.Add(c)
	f.SetSizeRequest(VIEW_WIDTH, util.GetIntPortion(util.HEIGHT, 1 - (2 * util.MARGIN)))
	f.SetMarginTop(util.GetIntPortion(util.HEIGHT, util.MARGIN))
	f.SetMarginBottom(util.GetIntPortion(util.HEIGHT, util.MARGIN))
	return &StatusView{
		Frame: f,
		Box: c,
		Mutex: &sync.RWMutex{},
	}
}

/*
Adds a @StatusMember to the StatusView. This will queue a Write operation
*/
func (s *StatusView) AddMember(name string) {
	m := StatusMemberNew(name)
	s.Mutex.Lock()
	s.Users = append(s.Users, m)
	glib.IdleAdd(func() {s.Box.PackStart(m, true, true, 0); m.ShowAll()})
	s.Mutex.Unlock()
}

/*
Removes a @StatusMember from the StatusView. This will queue a Read operation
*/
func (s *StatusView) RemoveMember(name string) {
	s.Mutex.RLock()
	for index, user := range(s.Users) {
		if user.Name == name {
			user.Destroy()
			s.Users = append(s.Users[:index], s.Users[index+1:]...)
			break
		}
	}
	s.Mutex.RUnlock()
}