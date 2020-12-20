package conn

import (
	"bufio"
	"log"
	"net"

	"github.com/client/ui"
)

/*
Connect dials to the server, establishes a connection, keeps a routine to listen
for incomming commands from the server and returns an error if there is one, else
it returns nil
*/
func Connect(window *ui.MainWindow) error {

	conn, err := net.Dial("tcp", ":9000")

	if err != nil {
		return err
	}

	go listen(conn, window)

	err = SendCommand(conn, CMD_JOIN_REQUEST, []byte(window.Username))

	p := window.Chat.Input.Entry

	p.Connect("activate", func() {
		t, _ := p.GetText()
		p.SetText("")
		SendCommand(conn, CMD_MESSAGE, []byte(t))
	})

	return err
}

func listen(conn net.Conn, window *ui.MainWindow) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		c := RetrieveCommand(conn)
		var err error
		switch c.id {
		case CMD_JOIN_NOTIF:
			window.Status.AddMember(string(c.data))
		case CMD_EXIT_NOTIF:
			window.Status.RemoveMember(string(c.data))
		case CMD_MESSAGE:
			err = window.Chat.Area.SendMessage(BytesToOwnerMessage(c.data))
		case CMD_PING:
			err = SendCommand(conn, CMD_PING_RESPONSE, []byte{})
		}
		if err != nil {
			log.Fatalln("Error while recieving external command: " + err.Error())
		}
	}

}
